jQuery(document).ready(function($){var i=$(".product form.variations_form"),n=i.find(".single_variation_wrap");i.on("reset_data",function(){i.find(".single_variation_wrap_kad").find(".quantity").hide(),i.find(".single_variation .price").hide()}),i.on("woocommerce_variation_has_changed",function(){$(".kad-select").trigger("update")}),n.on("hide_variation",function(){$(this).css("height","auto")})});;jQuery(function($){$("div.quantity:not(.buttons_added), td.quantity:not(.buttons_added)").addClass("buttons_added").append('<input type="button" value="+" class="plus" />').prepend('<input type="button" value="-" class="minus" />'),$(document).on("click",".plus, .minus",function(){var t=$(this).closest(".quantity").find(".qty"),a=parseFloat(t.val()),n=parseFloat(t.attr("max")),s=parseFloat(t.attr("min")),e=t.attr("step");a&&""!==a&&"NaN"!==a||(a=0),(""===n||"NaN"===n)&&(n=""),(""===s||"NaN"===s)&&(s=0),("any"===e||""===e||void 0===e||"NaN"===parseFloat(e))&&(e=1),$(this).is(".plus")?t.val(n&&(n==a||a>n)?n:a+parseFloat(e)):s&&(s==a||s>a)?t.val(s):a>0&&t.val(a-parseFloat(e)),t.trigger("change")})});