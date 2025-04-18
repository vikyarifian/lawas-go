// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"lawas-go/dto"
	"lawas-go/models"
)

func HtmlSell(token dto.Token, categories []models.Category, currencies []models.Currency) templ.Component {
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
		if token.IsAuth {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"modal fade\" id=\"sell-modal\" tabindex=\"-1\" role=\"dialog\" aria-hidden=\"true\"><div class=\"modal-dialog modal-dialog-centered\" role=\"document\"><div class=\"modal-content\"><div class=\"modal-body\"><button type=\"button\" class=\"close\" data-dismiss=\"modal\" aria-label=\"Close\"><span aria-hidden=\"true\"><i class=\"icon-close\"></i></span></button><div class=\"form-box\"><div class=\"form-tab\"><ul class=\"nav nav-pills nav-fill nav-border-anim\" role=\"tablist\"><li class=\"nav-item\"><a class=\"nav-link active\" id=\"sell-tab\" data-toggle=\"tab\" href=\"#sell\" role=\"tab\" aria-controls=\"sell\" aria-selected=\"true\">Start Offer</a></li></ul><div class=\"tab-content\" id=\"tab-content-5\"><div class=\"tab-pane fade show active\" id=\"sell\" role=\"tabpanel\" aria-labelledby=\"sell-tab\"><form id=\"sell-form\" enctype=\"multipart/form-data\" hx-disabled-elt=\"#sell-button\" hx-post=\"/sell\" hx-swap=\"none\" hx-on::before-request=\"document.getElementById(&#39;error_sell&#39;).innerHTML = &#39;&#39;\" hx-target-400=\"#error_sell\" hx-indicator=\".htmx-loader\"><div class=\"form-group\"><label for=\"sell-name\">Title *</label> <input type=\"text\" class=\"form-control\" id=\"sell-name\" name=\"name\" required></div><!-- End .form-group --><div class=\"form-group\"><label for=\"sell-description\">Description *</label> <textarea class=\"form-control\" id=\"sell-description\" name=\"description\" required style=\"white-space: normal|nowrap|pre|pre-line|pre-wrap|initial|inherit;\"></textarea></div><!-- End .form-group --><div class=\"form-group\"><label for=\"sell-category\">Category *</label> <select class=\"form-control\" id=\"sell-category\" name=\"category_id\" required>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, cate := range categories {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<option name=\"category_id\" value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var2 string
				templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(cate.ID)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/sell.templ`, Line: 51, Col: 97}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(cate.Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/sell.templ`, Line: 51, Col: 109}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "</option>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</select></div><!-- End .form-group --><div class=\"form-group\"><label for=\"sell-brand\">Brand </label> <input type=\"text\" class=\"form-control\" id=\"sell-brand\" name=\"brand\"></div><!-- End .form-group --><div class=\"form-group\"><label for=\"sell-condition\">Condition *</label> <select class=\"form-condition\" id=\"sell-condition\" name=\"condition\" required><option name=\"condition\" value=\"1\">New</option> <option name=\"condition\" value=\"2\">Used</option></select></div><!-- End .form-group --><div class=\"form-group\"><label for=\"sell-brand\">Duration *</label> <select class=\"form-duration\" id=\"sell-duration\" name=\"duration\" required><option name=\"condition\" value=\"3\">3 days</option> <option name=\"condition\" value=\"5\">5 days</option> <option name=\"condition\" value=\"7\">7 days</option> <option name=\"condition\" value=\"10\">10 days</option></select></div><!-- End .form-group --><div class=\"form-group\"><label for=\"sell-currency\">Currnecy *</label> <select class=\"form-control\" id=\"sell-currency\" name=\"currency_id\" required>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, curr := range currencies {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<option name=\"currency_id\" value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(curr.ID)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/sell.templ`, Line: 83, Col: 97}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(curr.Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/sell.templ`, Line: 83, Col: 109}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "</option>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "</select></div><!-- End .form-group --><div class=\"form-group\"><label for=\"sell-price\">Open Bid (Start Price) *</label> <input type=\"number\" class=\"form-control\" id=\"price\" name=\"open_bid\" required></div><!-- End .form-group --><div class=\"form-group\"><label for=\"sell-photo\">Photo *</label> <input type=\"file\" class=\"form-control\" id=\"photo\" name=\"photo\" required></div><!-- End .form-group --><span id=\"error_sell\"></span> <span id=\"success_sell\"></span><div class=\"form-footer\"><button id=\"sell-button\" type=\"submit\" class=\"btn btn-outline-primary-2\"><span>SAVE</span> <i class=\"icon-long-arrow-right\"></i></button></div><!-- End .form-footer --></form></div><!-- .End .tab-pane --></div><!-- End .tab-content --></div><!-- End .form-tab --></div><!-- End .form-box --></div><!-- End .modal-body --></div><!-- End .modal-content --></div><!-- End .modal-dialog --></div><!-- End .modal -->")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
