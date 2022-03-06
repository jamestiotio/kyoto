package kyoto

// Deprecated
var ssaclient = "<script>(()=>{var q=11;function ie(e,r){var t=r.attributes,n,a,i,d,c;if(!(r.nodeType===q||e.nodeType===q)){for(var v=t.length-1;v>=0;v--)n=t[v],a=n.name,i=n.namespaceURI,d=n.value,i?(a=n.localName||a,c=e.getAttributeNS(i,a),c!==d&&(n.prefix===\"xmlns\"&&(a=n.name),e.setAttributeNS(i,a,d))):(c=e.getAttribute(a),c!==d&&e.setAttribute(a,d));for(var h=e.attributes,A=h.length-1;A>=0;A--)n=h[A],a=n.name,i=n.namespaceURI,i?(a=n.localName||a,r.hasAttributeNS(i,a)||e.removeAttributeNS(i,a)):r.hasAttribute(a)||e.removeAttribute(a)}}var B,le=\"http://www.w3.org/1999/xhtml\",p=typeof document==\"undefined\"?void 0:document,oe=!!p&&\"content\"in p.createElement(\"template\"),se=!!p&&p.createRange&&\"createContextualFragment\"in p.createRange();function de(e){var r=p.createElement(\"template\");return r.innerHTML=e,r.content.childNodes[0]}function ue(e){B||(B=p.createRange(),B.selectNode(p.body));var r=B.createContextualFragment(e);return r.childNodes[0]}function fe(e){var r=p.createElement(\"body\");return r.innerHTML=e,r.childNodes[0]}function ce(e){return e=e.trim(),oe?de(e):se?ue(e):fe(e)}function R(e,r){var t=e.nodeName,n=r.nodeName,a,i;return t===n?!0:(a=t.charCodeAt(0),i=n.charCodeAt(0),a<=90&&i>=97?t===n.toUpperCase():i<=90&&a>=97?n===t.toUpperCase():!1)}function pe(e,r){return!r||r===le?p.createElement(e):p.createElementNS(r,e)}function ve(e,r){for(var t=e.firstChild;t;){var n=t.nextSibling;r.appendChild(t),t=n}return r}function G(e,r,t){e[t]!==r[t]&&(e[t]=r[t],e[t]?e.setAttribute(t,\"\"):e.removeAttribute(t))}var z={OPTION:function(e,r){var t=e.parentNode;if(t){var n=t.nodeName.toUpperCase();n===\"OPTGROUP\"&&(t=t.parentNode,n=t&&t.nodeName.toUpperCase()),n===\"SELECT\"&&!t.hasAttribute(\"multiple\")&&(e.hasAttribute(\"selected\")&&!r.selected&&(e.setAttribute(\"selected\",\"selected\"),e.removeAttribute(\"selected\")),t.selectedIndex=-1)}G(e,r,\"selected\")},INPUT:function(e,r){G(e,r,\"checked\"),G(e,r,\"disabled\"),e.value!==r.value&&(e.value=r.value),r.hasAttribute(\"value\")||e.removeAttribute(\"value\")},TEXTAREA:function(e,r){var t=r.value;e.value!==t&&(e.value=t);var n=e.firstChild;if(n){var a=n.nodeValue;if(a==t||!t&&a==e.placeholder)return;n.nodeValue=t}},SELECT:function(e,r){if(!r.hasAttribute(\"multiple\")){for(var t=-1,n=0,a=e.firstChild,i,d;a;)if(d=a.nodeName&&a.nodeName.toUpperCase(),d===\"OPTGROUP\")i=a,a=i.firstChild;else{if(d===\"OPTION\"){if(a.hasAttribute(\"selected\")){t=n;break}n++}a=a.nextSibling,!a&&i&&(a=i.nextSibling,i=null)}e.selectedIndex=t}}},y=1,he=11,k=3,W=8;function N(){}function ge(e){if(e)return e.getAttribute&&e.getAttribute(\"id\")||e.id}function Ae(e){return function(t,n,a){if(a||(a={}),typeof n==\"string\")if(t.nodeName===\"#document\"||t.nodeName===\"HTML\"||t.nodeName===\"BODY\"){var i=n;n=p.createElement(\"html\"),n.innerHTML=i}else n=ce(n);var d=a.getNodeKey||ge,c=a.onBeforeNodeAdded||N,v=a.onNodeAdded||N,h=a.onBeforeElUpdated||N,A=a.onElUpdated||N,ee=a.onBeforeNodeDiscarded||N,E=a.onNodeDiscarded||N,te=a.onBeforeElChildrenUpdated||N,P=a.childrenOnly===!0,w=Object.create(null),S=[];function O(s){S.push(s)}function K(s,o){if(s.nodeType===y)for(var l=s.firstChild;l;){var u=void 0;o&&(u=d(l))?O(u):(E(l),l.firstChild&&K(l,o)),l=l.nextSibling}}function H(s,o,l){ee(s)!==!1&&(o&&o.removeChild(s),E(s),K(s,l))}function J(s){if(s.nodeType===y||s.nodeType===he)for(var o=s.firstChild;o;){var l=d(o);l&&(w[l]=o),J(o),o=o.nextSibling}}J(t);function I(s){v(s);for(var o=s.firstChild;o;){var l=o.nextSibling,u=d(o);if(u){var g=w[u];g&&R(o,g)?(o.parentNode.replaceChild(g,o),D(g,o)):I(o)}else I(o);o=l}}function re(s,o,l){for(;o;){var u=o.nextSibling;(l=d(o))?O(l):H(o,s,!0),o=u}}function D(s,o,l){var u=d(o);u&&delete w[u],!(!l&&(h(s,o)===!1||(e(s,o),A(s),te(s,o)===!1)))&&(s.nodeName!==\"TEXTAREA\"?ne(s,o):z.TEXTAREA(s,o))}function ne(s,o){var l=o.firstChild,u=s.firstChild,g,b,L,x,T;e:for(;l;){for(x=l.nextSibling,g=d(l);u;){if(L=u.nextSibling,l.isSameNode&&l.isSameNode(u)){l=x,u=L;continue e}b=d(u);var _=u.nodeType,m=void 0;if(_===l.nodeType&&(_===y?(g?g!==b&&((T=w[g])?L===T?m=!1:(s.insertBefore(T,u),b?O(b):H(u,s,!0),u=T):m=!1):b&&(m=!1),m=m!==!1&&R(u,l),m&&D(u,l)):(_===k||_==W)&&(m=!0,u.nodeValue!==l.nodeValue&&(u.nodeValue=l.nodeValue))),m){l=x,u=L;continue e}b?O(b):H(u,s,!0),u=L}if(g&&(T=w[g])&&R(T,l))s.appendChild(T),D(T,l);else{var $=c(l);$!==!1&&($&&(l=$),l.actualize&&(l=l.actualize(s.ownerDocument||p)),s.appendChild(l),I(l))}l=x,u=L}re(s,u,b);var j=z[s.nodeName];j&&j(s,o)}var f=t,U=f.nodeType,X=n.nodeType;if(!P){if(U===y)X===y?R(t,n)||(E(t),f=ve(t,pe(n.nodeName,n.namespaceURI))):f=n;else if(U===k||U===W){if(X===U)return f.nodeValue!==n.nodeValue&&(f.nodeValue=n.nodeValue),f;f=n}}if(f===n)E(t);else{if(n.isSameNode&&n.isSameNode(f))return;if(D(f,n,P),S)for(var F=0,ae=S.length;F<ae;F++){var V=w[S[F]];V&&H(V,V.parentNode,!1)}}return!P&&f!==t&&t.parentNode&&(f.actualize&&(f=f.actualize(t.ownerDocument||p)),t.parentNode.replaceChild(f,t)),f}}var be=Ae(ie),Y=be;function C(e){let r=e.starter;if(e.id){let t=document.getElementById(e.id);if(!t)throw new Error(`Error while locating root with id: can't find direct with ${e}`);r=t}else{let t=0;for(;;){if(!r.parentElement)throw new Error(`Error while locating root: can't find parent with ${e}`);if(!r.getAttribute(\"state\"))r=r.parentElement;else if(e.depth&&t!=e.depth)r=r.parentElement,t++;else break}}return r}function Q(e){return e=unescape(encodeURIComponent(e)),e=btoa(e),e=e.replaceAll(\"/\",\"-\"),e}function Te(e){return e=e.replaceAll(\"-\",\"/\"),e=atob(e),e=decodeURIComponent(escape(e)),e}function me(e){return e.includes(\":\")&&(e=e.split(\":\")[1]),e.includes(\"$\")&&(e=e.replaceAll(\"$\",\"\")),e}function Ne(e){e.querySelectorAll(\"[ssa\\\\:oncall\\\\.display]\").forEach(t=>{let n=t.getAttribute(\"ssa:oncall.display\");n!=\"\"&&t.setAttribute(\"style\",\"display: \"+n)})}function we(){document.querySelectorAll(\"[ssa\\\\:onload]\").forEach(e=>{let r=e.getAttribute(\"ssa:onload\");r&&r!=\"\"&&M(e,r)})}function Le(){document.querySelectorAll(\"[ssa\\\\:poll]\").forEach(e=>{let r=e.getAttribute(\"ssa:poll\")||\"\",t=e.getAttribute(\"ssa:poll.interval\");r&&r!=\"\"&&t&&t!=\"\"&&setInterval(()=>{M(e,r)},parseInt(t))})}function ye(){document.querySelectorAll(\"[ssa\\\\:onintersect]\").forEach(e=>{let r=e.getAttribute(\"ssa:onintersect\")||\"\",t=e.getAttribute(\"ssa:onintersect.threshold\")||\"1.0\";r!=\"\"&&new IntersectionObserver(a=>{a.forEach(i=>{i.intersectionRatio>=parseFloat(t)&&M(e,r,parseFloat(t))})},{threshold:parseFloat(t)}).observe(e)})}function Z(e,r,t){let n=new Array,a={};t&&(a=t),a.onBeforeElUpdated=function(i,d){if(i.getAttribute(\"ssa:morph.ignore.attr\")!=null){let c=i.getAttribute(\"ssa:morph.ignore.attr\");if(c)if(c==\"innerHTML\")d.innerHTML=i.innerHTML;else{let v=i.getAttribute(c);v&&d.setAttribute(c,v)}}return i.getAttribute(\"ssa:morph.ignore\")!=null?!1:i.getAttribute(\"ssa:morph.ignore.this\")!=null&&i!=e?(n.push({fromEl:i,toEl:d}),!1):!0},Y(e,r,a),n.length>0&&n.forEach(i=>{Z(i.fromEl,i.toEl,{childrenOnly:!0})})}function M(e,r,...t){return new Promise((n,a)=>{let i=C({starter:e,depth:r.split(\"\").filter(v=>v===\"$\").length,id:r.includes(\":\")?r.split(\":\")[0]:void 0});Ne(i);let d=ssapath;d+=`/${i.getAttribute(\"name\")}`,d+=`/${i.getAttribute(\"state\")||\"{}\"}`,d+=`/${me(r)}`,d+=`/${Q(JSON.stringify(t))}`;let c=new EventSource(d);c.onmessage=v=>{let h=v.data;if(!!h){if(h.startsWith(\"ssa:redirect=\")){let A=h.replace(\"ssa:redirect=\",\"\");window.location.href=A;return}if(i.getAttribute(\"ssa:render.mode\")==\"replace\"){i.outerHTML=h;return}try{Z(i,h)}catch(A){console.log(\"Fallback from morphdom to root.outerHTML due to error\",A),i.outerHTML=h}}},c.onerror=v=>{c.close(),n()}})}function Me(e,r){let t=C({starter:e,depth:r.split(\"\").filter(a=>a===\"$\").length,id:r.includes(\":\")?r.split(\":\")[0]:void 0});if(!t.getAttribute(\"state\"))throw new Error(\"Bind call error: component state is underfined\");let n=JSON.parse(Te(t.getAttribute(\"state\")));n[r]=e.value,t.setAttribute(\"state\",Q(JSON.stringify(n)))}function Ee(e,r){r.preventDefault();let t=C({starter:e});if(!t.getAttribute(\"state\"))throw new Error(\"Bind call error: component state is underfined\");let n=JSON.parse(atob(t.getAttribute(\"state\"))),a=new FormData(r.target),i=Object.fromEntries(a.entries());return Object.entries(i).forEach(d=>{n[d[0]]=d[1]}),t.setAttribute(\"state\",btoa(JSON.stringify(n).replaceAll(\"/\",\"-\"))),M(t,\"Submit\"),!1}window._LocaleRoot=C;window.Action=M;window.Bind=Me;window.FormSubmit=Ee;document.addEventListener(\"DOMContentLoaded\",we);document.addEventListener(\"DOMContentLoaded\",ye);document.addEventListener(\"DOMContentLoaded\",Le);})();</script>"
