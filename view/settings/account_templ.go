// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package settings

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"ftgodev-tut/models"
	"ftgodev-tut/view"
	"ftgodev-tut/view/layout"
	"ftgodev-tut/view/ui"
)

type ProfileParams struct {
	Username string
	Success  bool
}

type ProfileFormErrors struct {
	Username string
}

func Index(user models.AuthenticatedUser) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = layout.App(true).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"max-w-2xl w-full mx-auto mt-8\"><div><h1 class=\" text-lg font-semibold border-b border-app pb-2\">Profile</h1>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ProfileSettings(
			ProfileParams{Username: user.Account.Username},
			ProfileFormErrors{},
		).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"mt-10\"><h1 class=\" text-lg font-semibold border-b border-app pb-2\">Credits</h1><div class=\"sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0 items-center mt-8\"><dt class=\"\">Credits</dt><dd class=\"sm:col-span-1 sm:mt-0\"><span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(view.String(user.Account.Credits))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/settings/account.templ`, Line: 34, Col: 50}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></dd><dd class=\"sm:col-span-1 sm:mt-0\"><a href=\"/buy-credits\" class=\"btn btn-outline\"><i class=\"fa-solid fa-money-bill-transfer\"></i> Buy Credits</a></dd></div></div><div class=\"mt-10\"><h1 class=\" text-lg font-semibold border-b border-app pb-2\">Change Password</h1><div class=\"sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0 items-center mt-8\"><dt class=\"\">Reset Password</dt><dd class=\"\"><button hx-post=\"/auth/reset-password\" hx-swap=\"outerHTML\" class=\"btn btn-primary\">Reset Password</button></dd></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func ProfileSettings(params ProfileParams, errors ProfileFormErrors) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form hx-put=\"/settings/account/profile\" hx-swap=\"outerHTML\"><div class=\"sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0 items-center mt-8\"><dt class=\"\">User Name</dt><dd class=\"sm:col-span-2 sm:mt-0\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if params.Success {
			templ_7745c5c3_Err = ui.Toast("Profile Saved").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input class=\"input input-bordered w-full max-w-sm\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(params.Username)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/settings/account.templ`, Line: 75, Col: 34}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" name=\"username\"> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(errors.Username) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"text-sm text-error\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(errors.Username)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/settings/account.templ`, Line: 78, Col: 60}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</dd><dt></dt><dd><button type=\"submit\" class=\"btn btn-primary\">save</button></dd></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

type PasswordFormParams struct {
	OldPassword     string
	NewPassword     string
	ConfirmPassword string
}

type PasswordFormErrors struct {
	OldPassword     string
	NewPassword     string
	ConfirmPassword string
}

func PasswordForm(params PasswordFormParams, errors PasswordFormErrors) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form hx-put=\"settings/account/password\" hx-swap=\"outerHTML\"><div class=\"sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0 items-center mt-8\"><dt class=\"\">Old Password</dt><dd class=\"sm:col-span-2 sm:mt-0\"><input class=\"input input-bordered w-full max-w-sm\"></dd><dt class=\"\">New Password</dt><dd class=\"sm:col-span-2 sm:mt-0\"><input class=\"input input-bordered w-full max-w-sm\"></dd><dt class=\"\">Can you type it again?</dt><dd class=\"sm:col-span-2 sm:mt-0\"><input class=\"input input-bordered w-full max-w-sm\"></dd><dt></dt><dd><button class=\"btn btn-primary\">update password</button></dd></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
