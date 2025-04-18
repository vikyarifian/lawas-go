// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"lawas-go/components"
	"lawas-go/dto"
)

func About(token dto.Token, isLoggedIn bool) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<nav aria-label=\"breadcrumb\" class=\"breadcrumb-nav mb-3\"><div class=\"container\"><ol class=\"breadcrumb\"><li class=\"breadcrumb-item\"><a href=\"/\">Home</a></li><li class=\"breadcrumb-item\"><a href=\"#\">Pages</a></li><li class=\"breadcrumb-item active\" aria-current=\"page\">About us</li></ol></div><!-- End .container --></nav><!-- End .breadcrumb-nav --> <div class=\"container\"><div class=\"page-header page-header-big text-center\" style=\"background-image: url(&#39;assets/images/slider/slide-1.jpg&#39;)\"><h1 class=\"page-title text-white\">About us<span class=\"text-white\">who we are</span></h1></div><!-- End .page-header --></div><!-- End .container -->  <div class=\"container\"><div class=\"row\"><div class=\"col-lg-5 mb-3 mb-lg-0\"><h2 class=\"title\">Who We Are</h2><!-- End .title --><p class=\"lead text-primary mb-3\">The ultimate online market platform where buyers and sellers connect for the best deals! </p><!-- End .lead text-primary --><p class=\"mb-2\">Whether you're looking for rare collectibles, the latest gadgets, or unique treasures, LAWAS brings the excitement of auctions straight to your fingertips. </p></div><!-- End .col-lg-5 --><div class=\"col-lg-6 offset-lg-1\"><div class=\"about-images\"><img src=\"assets/images/about/img-3.png\" style=\"border: 0cap;\" alt=\"\" class=\"about-img-front\"> <img src=\"assets/images/about/img-4.png\" style=\"border: 0cap;\" alt=\"\" class=\"about-img-back\"></div><!-- End .about-images --></div><!-- End .col-lg-6 --></div><!-- End .row --></div><!-- End .container --> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = components.Layout(token, isLoggedIn, "About").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
