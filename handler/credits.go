package handler

import (
  "fmt"
  "net/http"
  "os"

  "ftgodev-tut/view/credits"
  "ftgodev-tut/db"

  "github.com/go-chi/chi/v5"
  "github.com/stripe/stripe-go/v76"
  "github.com/stripe/stripe-go/v76/checkout/session"
)

func HandleCreditsIndex(w http.ResponseWriter, r *http.Request) error {
  return render(r, w, credits.Index()) 

}

func HandleStripeCheckoutCreate(w http.ResponseWriter, r *http.Request) error {
  stripe.Key = os.Getenv("STRIPE_API_KEY")
  checkoutParams := &stripe.CheckoutSessionParams{
    //SuccessURL: stripe.String(os.Getenv("STRIPE_SUCESS_URL")),
    SuccessURL: stripe.String("http://localhost:7331/checkout/success/{CHECKOUT_SESSION_ID}"),
    CancelURL: stripe.String(os.Getenv("STRIPE_CANCEL_URL")),
    Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
    LineItems: []*stripe.CheckoutSessionLineItemParams{
      {
        Price: stripe.String(chi.URLParam(r, "priceID")),
        Quantity: stripe.Int64(1),
      },
    },
  }

  sesh, err := session.New(checkoutParams)
  fmt.Println("credits handler")
  fmt.Println(err)
  fmt.Println(sesh.URL)

  if err != nil {
    return err
  }

  hxRedirect(w, r, sesh.URL)
  return nil
}

func HandleStripeCheckoutSuccess(w http.ResponseWriter, r *http.Request) error {
  user := getAuthenticatedUser(r)
  sessionID := chi.URLParam(r, "sessionID")
  stripe.Key = os.Getenv("STRIPE_API_KEY")
  sesh, err := session.Get(sessionID, nil)
  if err != nil {
    return err
  }

  lineItemParams := stripe.CheckoutSessionListLineItemsParams{}
  lineItemParams.Session = stripe.String(sesh.ID)
  iter := session.ListLineItems(&lineItemParams)
  iter.Next() // we only care about the first item
  item := iter.LineItem()
  priceID := item.Price.ID

  switch priceID {
  case os.Getenv("ONEHUNDO_PRICE_ID"):
    user.Account.Credits += 100
  case os.Getenv("TWOFIFTY_PRICE_ID"):
    user.Account.Credits += 250
  default: 
    return fmt.Errorf("invalid price id")
  }

  if err := db.UpdateAccount(&user.Account); err != nil {
    return err
  }

  // ideally we would save their callback in local storage, cookie
  http.Redirect(w, r, "/generate", http.StatusSeeOther)

  return nil
}

func HandleStripeCheckoutCancel(w http.ResponseWriter, r *http.Request) error {
  return nil
}
