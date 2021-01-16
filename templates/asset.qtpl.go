// Code generated by qtc from "asset.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/asset.qtpl:1
package templates

//line templates/asset.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/asset.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/asset.qtpl:1
func StreamDefaultCSS(qw422016 *qt422016.Writer) {
//line templates/asset.qtpl:1
	qw422016.N().S(`
`)
//line templates/asset.qtpl:2
	qw422016.N().S(`/* Layout stuff */
@media screen and (min-width: 800px) {
	main {padding:1rem 2rem; margin: 0 auto; width: 800px;}
}
@media screen and (max-width: 800px) {
	main {padding: 1rem; margin: 0; width: 100%;}
}
*, *::before, *::after {box-sizing: border-box;}
html {height:100%; padding:0;
background-image: url("data:image/svg+xml,%3Csvg width='42' height='44' viewBox='0 0 42 44' xmlns='http://www.w3.org/2000/svg'%3E%3Cg id='Page-1' fill='none' fill-rule='evenodd'%3E%3Cg id='brick-wall' fill='%23bbbbbb' fill-opacity='0.4'%3E%3Cpath d='M0 0h42v44H0V0zm1 1h40v20H1V1zM0 23h20v20H0V23zm22 0h20v20H22V23z'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E");} /* heropatterns.com */
body {height:100%; margin:0; font-size:16px; font-family: 'PT Sans', 'Liberation Sans', sans-serif;}
main {border-radius: 0 0 .25rem .25rem; }
main > form {margin-bottom:1rem;}
textarea {font-size:16px; font-family: 'PT Sans', 'Liberation Sans', sans-serif;}
.edit_no-preview {height:100%;}
.edit_with-preview .edit-form textarea { min-height: 500px; }
.edit__preview { border: 2px dashed #ddd; }
.edit-form {height:90%;}
.edit-form textarea {width:100%;height:90%;}
.edit-form__save { font-weight: bold; }
.icon {margin-right: .25rem; vertical-align: bottom; }

main h1:not(.navi-title) {font-size:1.7rem;}
blockquote { margin-left: 0; padding-left: 1rem; }
.wikilink__destination-type {display: inline; margin: 0 .25rem; vertical-align: sub; }

article { overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; line-height: 150%; }
article h1, article h2, article h3, article h4, article h5, article h6 { margin: 1.5rem 0 0 0; }
article p { margin: .5rem 0; }
article ul, ol { padding-left: 1.5rem; margin: .5rem 0; }
article code { padding: .1rem .3rem; border-radius: .25rem; font-size: 90%; }
article pre.codeblock { padding:.5rem; white-space: pre-wrap; border-radius: .25rem;}
.codeblock code {padding:0; font-size:15px;}
.transclusion { border-radius: .25rem; }
.transclusion__content > *:not(.binary-container) {margin: 0.5rem; }
.transclusion__link {display: block; text-align: right; font-style: italic; margin-top: .5rem; margin-right: .25rem; text-decoration: none;}
.transclusion__link::before {content: "⇐ ";}

/* Derived from https://commons.wikimedia.org/wiki/File:U%2B21D2.svg */
.launchpad__entry { list-style-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' version='1.0' width='25' height='12'%3E%3Cg transform='scale(0.7,0.8) translate(-613.21429,-421)'%3E%3Cpath d='M 638.06773,429.49751 L 631.01022,436.87675 L 630.1898,436.02774 L 632.416,433.30375 L 613.46876,433.30375 L 613.46876,431.66382 L 633.82089,431.66382 L 635.57789,429.5261 L 633.79229,427.35979 L 613.46876,427.35979 L 613.46876,425.71985 L 632.416,425.71985 L 630.1898,422.99587 L 631.01022,422.08788 L 638.06773,429.49751 z '/%3E%3C/g%3E%3C/svg%3E"); }

.binary-container_with-img img,
.binary-container_with-video video,
.binary-container_with-audio audio {width: 100%}

.navi-title { padding-bottom: .5rem; margin-bottom: .25rem; }
.navi-title a {text-decoration:none; }
.navi-title__separator { margin: 0 .25rem; }
.navi-title__colon { margin-right: .5rem; }
.upload-amnt { clear: both; padding: .5rem; border-radius: .25rem; }
aside { clear: both; }

.img-gallery { text-align: center; margin-top: .25rem; margin-bottom: .25rem; }
.img-gallery_many-images { border-radius: .25rem; padding: .5rem; }
.img-gallery img { max-width: 100%; max-height: 50vh; }
figure { margin: 0; }
figcaption { padding-bottom: .5rem; }

nav ul {display:flex; padding-left:0; flex-wrap:wrap; margin-top:0;}
nav ul li {list-style-type:none;margin-right:1rem;}

#new-name {width:100%;}
.navlinks b {text-decoration: underline; }
.navlinks__user {font-style:italic;}

.rc-entry { display: grid; list-style-type: none; padding: .25rem; grid-template-columns: 1fr 1fr; }
.rc-entry__time { font-style: italic; }
.rc-entry__hash { font-style: italic; text-align: right; }
.rc-entry__links { grid-column: 1 / span 2; }
.rc-entry__author { font-style: italic; }

.prevnext__el { display: block-inline; min-width: 40%; padding: .5rem; margin-bottom: .25rem; text-decoration: none; border-radius: .25rem; }
.prevnext__prev { float: left; }
.prevnext__next { float: right; text-align: right; }

.page-separator { clear: both; }
.history__entries { background-color: #eee; margin: 0; padding: 0; border-radius: .25rem; }
.history__month-anchor { text-decoration: none; color: black; }
.history__entry { list-style-type: none; padding: .25rem; }
.history-entry { padding: .25rem; }
.history-entry__time { font-weight: bold; }
.history-entry__author { font-style: italic; }

table { border: #ddd 1px solid; border-radius: .25rem; min-width: 4rem; }
td { padding: .25rem; }
caption { caption-side: top; font-size: small; }

/* Color stuff */
svg path { fill: currentColor !important; }
/* Lighter stuff #eee */
article code,
article .codeblock,
.transclusion,
.img-gallery_many-images,
.rc-entry,
.prevnext__el,
table { background-color: #eee; }

/* Other stuff */
html { background-color: #ddd; }

main { background-color: white; box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.2); }
blockquote { border-left: 4px black solid; }
.wikilink_new {color:#a55858;}
.transclusion code, .transclusion .codeblock {background-color:#ddd;}
.transclusion__link { color: black; }
.wikilink_new:visited {color:#a55858;}
.navi-title { border-bottom: #eee 1px solid; }
.upload-amnt { border: #eee 1px solid; }
td { border: #ddd 1px solid; }

/* Dark theme! */
@media (prefers-color-scheme: dark) {
html { background-color: #121212; color: #ddd; }
main,  article { background-color: #343434; }
blockquote { border-left: 4px #ddd solid; }
a, .wikilink_external { color: #f1fa8c; }
a:visited, .wikilink_external:visited { color: #ffb86c; }
.wikilink_new, .wikilink_new:visited { color: #dd4444; }
.prevnext__el, .prevnext__el:visited { color: #ddd; }

.transclusion .transclusion__link { color: #ddd; }
article code, 
article .codeblock, 
.transclusion,
.img-gallery_many-images,
.rc-entry,
.history__entry, 
.prevnext__el, 
.upload-amnt, 
textarea,
table { border: 0; background-color: #444444; color: #ddd; }
.transclusion code,
.transclusion .codeblock { background-color: #454545; }
}

`)
//line templates/asset.qtpl:2
	qw422016.N().S(`
`)
//line templates/asset.qtpl:3
}

//line templates/asset.qtpl:3
func WriteDefaultCSS(qq422016 qtio422016.Writer) {
//line templates/asset.qtpl:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/asset.qtpl:3
	StreamDefaultCSS(qw422016)
//line templates/asset.qtpl:3
	qt422016.ReleaseWriter(qw422016)
//line templates/asset.qtpl:3
}

//line templates/asset.qtpl:3
func DefaultCSS() string {
//line templates/asset.qtpl:3
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/asset.qtpl:3
	WriteDefaultCSS(qb422016)
//line templates/asset.qtpl:3
	qs422016 := string(qb422016.B)
//line templates/asset.qtpl:3
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/asset.qtpl:3
	return qs422016
//line templates/asset.qtpl:3
}

// Next three are from https://remixicon.com/

//line templates/asset.qtpl:6
func StreamIconHTTP(qw422016 *qt422016.Writer) {
//line templates/asset.qtpl:6
	qw422016.N().S(`
`)
//line templates/asset.qtpl:7
	qw422016.N().S(`<svg xmlns="http://www.w3.org/2000/svg" color="currentColor" viewBox="0 0 24 24" width="24" height="24"><path fill="none" d="M0 0h24v24H0z"/><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10-4.477 10-10 10zm-2.29-2.333A17.9 17.9 0 0 1 8.027 13H4.062a8.008 8.008 0 0 0 5.648 6.667zM10.03 13c.151 2.439.848 4.73 1.97 6.752A15.905 15.905 0 0 0 13.97 13h-3.94zm9.908 0h-3.965a17.9 17.9 0 0 1-1.683 6.667A8.008 8.008 0 0 0 19.938 13zM4.062 11h3.965A17.9 17.9 0 0 1 9.71 4.333 8.008 8.008 0 0 0 4.062 11zm5.969 0h3.938A15.905 15.905 0 0 0 12 4.248 15.905 15.905 0 0 0 10.03 11zm4.259-6.667A17.9 17.9 0 0 1 15.973 11h3.965a8.008 8.008 0 0 0-5.648-6.667z"/></svg>
`)
//line templates/asset.qtpl:7
	qw422016.N().S(`
`)
//line templates/asset.qtpl:8
}

//line templates/asset.qtpl:8
func WriteIconHTTP(qq422016 qtio422016.Writer) {
//line templates/asset.qtpl:8
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/asset.qtpl:8
	StreamIconHTTP(qw422016)
//line templates/asset.qtpl:8
	qt422016.ReleaseWriter(qw422016)
//line templates/asset.qtpl:8
}

//line templates/asset.qtpl:8
func IconHTTP() string {
//line templates/asset.qtpl:8
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/asset.qtpl:8
	WriteIconHTTP(qb422016)
//line templates/asset.qtpl:8
	qs422016 := string(qb422016.B)
//line templates/asset.qtpl:8
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/asset.qtpl:8
	return qs422016
//line templates/asset.qtpl:8
}

//line templates/asset.qtpl:10
func StreamIconGemini(qw422016 *qt422016.Writer) {
//line templates/asset.qtpl:10
	qw422016.N().S(`
`)
//line templates/asset.qtpl:11
	qw422016.N().S(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"><path fill="none" d="M0 0h24v24H0z"/><path d="M15.502 20A6.523 6.523 0 0 1 12 23.502 6.523 6.523 0 0 1 8.498 20h2.26c.326.489.747.912 1.242 1.243.495-.33.916-.754 1.243-1.243h2.259zM18 14.805l2 2.268V19H4v-1.927l2-2.268V9c0-3.483 2.504-6.447 6-7.545C15.496 2.553 18 5.517 18 9v5.805zM17.27 17L16 15.56V9c0-2.318-1.57-4.43-4-5.42C9.57 4.57 8 6.681 8 9v6.56L6.73 17h10.54zM12 11a2 2 0 1 1 0-4 2 2 0 0 1 0 4z"/></svg>
`)
//line templates/asset.qtpl:11
	qw422016.N().S(`
`)
//line templates/asset.qtpl:12
}

//line templates/asset.qtpl:12
func WriteIconGemini(qq422016 qtio422016.Writer) {
//line templates/asset.qtpl:12
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/asset.qtpl:12
	StreamIconGemini(qw422016)
//line templates/asset.qtpl:12
	qt422016.ReleaseWriter(qw422016)
//line templates/asset.qtpl:12
}

//line templates/asset.qtpl:12
func IconGemini() string {
//line templates/asset.qtpl:12
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/asset.qtpl:12
	WriteIconGemini(qb422016)
//line templates/asset.qtpl:12
	qs422016 := string(qb422016.B)
//line templates/asset.qtpl:12
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/asset.qtpl:12
	return qs422016
//line templates/asset.qtpl:12
}

//line templates/asset.qtpl:14
func StreamIconMailto(qw422016 *qt422016.Writer) {
//line templates/asset.qtpl:14
	qw422016.N().S(`
`)
//line templates/asset.qtpl:15
	qw422016.N().S(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"><path fill="none" d="M0 0h24v24H0z"/><path d="M3 3h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1zm17 4.238l-7.928 7.1L4 7.216V19h16V7.238zM4.511 5l7.55 6.662L19.502 5H4.511z"/></svg>
`)
//line templates/asset.qtpl:15
	qw422016.N().S(`
`)
//line templates/asset.qtpl:16
}

//line templates/asset.qtpl:16
func WriteIconMailto(qq422016 qtio422016.Writer) {
//line templates/asset.qtpl:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/asset.qtpl:16
	StreamIconMailto(qw422016)
//line templates/asset.qtpl:16
	qt422016.ReleaseWriter(qw422016)
//line templates/asset.qtpl:16
}

//line templates/asset.qtpl:16
func IconMailto() string {
//line templates/asset.qtpl:16
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/asset.qtpl:16
	WriteIconMailto(qb422016)
//line templates/asset.qtpl:16
	qs422016 := string(qb422016.B)
//line templates/asset.qtpl:16
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/asset.qtpl:16
	return qs422016
//line templates/asset.qtpl:16
}

// This is a modified version of https://www.svgrepo.com/svg/232085/rat

//line templates/asset.qtpl:19
func StreamIconGopher(qw422016 *qt422016.Writer) {
//line templates/asset.qtpl:19
	qw422016.N().S(`
`)
//line templates/asset.qtpl:20
	qw422016.N().S(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" width="24" height="24">
<path d="M447.238,204.944v-70.459c0-8.836-7.164-16-16-16c-34.051,0-64.414,21.118-75.079,55.286
C226.094,41.594,0,133.882,0,319.435c0,0.071,0.01,0.14,0.011,0.21c0.116,44.591,36.423,80.833,81.04,80.833h171.203
c8.836,0,16-7.164,16-16c0-8.836-7.164-16-16-16H81.051c-21.441,0-39.7-13.836-46.351-33.044H496c8.836,0,16-7.164,16-16
C512,271.82,486.82,228.692,447.238,204.944z M415.238,153.216v37.805c-10.318-2.946-19.556-4.305-29.342-4.937
C390.355,168.611,402.006,157.881,415.238,153.216z M295.484,303.435L295.484,303.435c-7.562-41.495-43.948-73.062-87.593-73.062
c-8.836,0-16,7.164-16,16c0,8.836,7.164,16,16,16c25.909,0,47.826,17.364,54.76,41.062H32.722
c14.415-159.15,218.064-217.856,315.136-90.512c3.545,4.649,9.345,6.995,15.124,6.118
c55.425-8.382,107.014,29.269,115.759,84.394H295.484z"/>
<circle cx="415.238" cy="260.05" r="21.166"/>
</svg>
`)
//line templates/asset.qtpl:20
	qw422016.N().S(`
`)
//line templates/asset.qtpl:21
}

//line templates/asset.qtpl:21
func WriteIconGopher(qq422016 qtio422016.Writer) {
//line templates/asset.qtpl:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/asset.qtpl:21
	StreamIconGopher(qw422016)
//line templates/asset.qtpl:21
	qt422016.ReleaseWriter(qw422016)
//line templates/asset.qtpl:21
}

//line templates/asset.qtpl:21
func IconGopher() string {
//line templates/asset.qtpl:21
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/asset.qtpl:21
	WriteIconGopher(qb422016)
//line templates/asset.qtpl:21
	qs422016 := string(qb422016.B)
//line templates/asset.qtpl:21
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/asset.qtpl:21
	return qs422016
//line templates/asset.qtpl:21
}
