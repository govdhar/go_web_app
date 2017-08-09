<!DOCTYPE html>
<html>
<head lang='en'>
<meta charset='UTF-8'>
<title>Stories</title>
<meta name='description' content='descrip'>
<meta name='keywords' content='keywords'>
<meta name='author' content='author'>
<meta name='viewport' content='width=device-width, initial-scale=1'>
<link rel='stylesheet' href='../public/css/css_reset.css'>
<link rel='stylesheet' href='http://fonts.googleapis.com/css?family=Droid+Sans:400'>
<link rel='stylesheet' href='../public/css/app.css'>
<link rel='stylesheet' href='../public/css/flyout_button.css'>
<link rel='stylesheet' href='../public/css/flyout_menu.css'>
<link rel='shortcut icon' href='../public//img/favicon.ico'>
</head>

<body>

<!--
**********************************
SIDEBAR
**********************************
-->

<aside class="sidebar" id="sidebar">
<nav class="menu-items">
<ul>
<li><a href="#">Home</a></li>
<li><a href="#">Me</a></li>
<li><a href="#">Interactions</a></li>
<li><a href="#">Something else</a></li>
<li><a href="#">Sign Out</a></li>
</ul>
</nav>


<!--
**********************************
HAMBURGER
**********************************
-->
<div id="hamburger" class="hamburglar menu-open-button">

<div class="burger-icon">
<div class="burger-container">
<span class="burger-bun-top"></span>
<span class="burger-filling"></span>
<span class="burger-bun-bot"></span>
</div>
</div>

<!-- svg ring containter -->
<div class="burger-ring">
<svg class="svg-ring">
<path class="path" fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="4"
d="M 34 2 C 16.3 2 2 16.3 2 34 s 14.3 32 32 32 s 32 -14.3 32 -32 S 51.7 2 34 2"/>
</svg>
</div>
<!-- the masked path that animates the fill to the ring -->

<svg width="0" height="0">
<mask id="mask">
<path xmlns="http://www.w3.org/2000/svg" fill="none" stroke="#000" stroke-miterlimit="10"
stroke-width="4" d="M 34 2 c 11.6 0 21.8 6.2 27.4 15.5 c 2.9 4.8 5 16.5 -9.4 16.5 h -4"/>
</mask>
</svg>
<div class="path-burger">
<div class="animate-path">
<div class="path-rotation"></div>
</div>
</div>

</div>

<!--
**********************************
HAMBURGER END
**********************************
-->
</aside>

<!--
**********************************
SIDEBAR END
**********************************
-->
<article>
<blockquote cite="http://codepen.io/GoesToEleven/pen/QbbVRa">
<p>One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible
vermin. He lay on his armour-like back, and if he lifted his head a little he could see his brown belly,
slightly domed and divided by arches into stiff sections. The bedding was hardly able to cover it and seemed
ready to slide off any moment. His many legs, pitifully thin compared with the size of the rest of him, waved
about helplessly as he looked.</p>
<p><cite>&#8212; <a href="http://codepen.io/GoesToEleven/pen/QbbVRa" target="_blank">The Metamorphoses</a></cite></p>
</blockquote>

</article>

<script src='../public/scripts/flyout_menu.js' async></script>
</body>
</html>