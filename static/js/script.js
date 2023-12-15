$(window).on('load', function() {
	$('.loader .inner').fadeOut(500, function() {
		$('.loader').fadeOut(750);
	});

	$('.projects').isotope({
    	itemSelector: '.project-item',
    	filter: '*',
    	animationOptions: {
			duration: 1500,
			easing: 'linear',
			queue: false
    	}
    });
});

$(document).ready(function(){
	console.log("Start.");
	$.i18n().load( {
		'en': './static/js/jquery.i18n/language/en.json',
		'pl': './static/js/jquery.i18n/language/pl.json'
	} ).done(function() {
		console.log("Languages loaded.");
	});

	$('.languageSelector').click(function() {
		$.i18n().locale = $(this).data('locale');
		$('body').i18n();
	});

	$('#slides').superslides({
		play: 4500,
		animation: 'fade',
		animations_speed: 'slow',
		pagination: false
	});

	var typedTitleSubText = new Typed('.titlePage .sub', {
		strings: ['Software Developer', "Freelancer", "Owner of the ThinkingInCode"],
		typeSpeed: 150,
		backSpeed: 30,
		loop: true,
		startDelay: 1000,
		showCursor: false
	});

	$('.owl-carousel').owlCarousel({
	    loop:true,
	    responsive:{
	        0:{
	            items:1
	        },
	        480:{
	            items:2
	        },
	        768:{
	            items:3
	        },
	        938:{
	            items:5
	        }
	    }
	});

	
    var skillsTopOffset = $('.skillsSection').offset().top;

    $(window).scroll(function() {
    	if(window.pageYOffset > skillsTopOffset - $(window).height() + 200) {
    		$('.pieChart').easyPieChart( {
		        easing: 'easeInOut',
		        barColor: '#fff',
		        trackColor: false,
		        lineWidth: 5,
		        lineCap: 'butt',
		        scaleLength: 0,
		        size: 152,
		        onStep: function(from, to, percent) {
		           	$(this.el).find('.percent').text(Math.round(percent));
		            }
		    });
    	}
    });

    updateStatsSectionElements();

    var statsTopOffset = $('.statsSection').offset().top;
    var firstShow = true
    $(window).scroll(function() {
    	if(firstShow && window.pageYOffset > statsTopOffset - $(window).height() + 200) {
    		$('.counter').countTo();

    		firstShow = false;
    	}
    });

    $('[data-fancybox]').fancybox();

    $('#filters a').click(function (){
    	$('#filters .current').removeClass('current');
    	$(this).addClass("current");

    	var selector = $(this).attr('data-filter');

    	$('.projects').isotope({
	    	filter: selector,
	    	animationOptions: {
				duration: 1500,
				easing: 'linear',
				queue: false
	    	}
    	});

    	return false; // Do nothing else
    });

    $('.forSlide a').click(function(e) {
    	e.preventDefault();

    	var target = $(this).attr("href");
    	var targetPosition = $(target).offset().top;

    	$('html, body').animate( { scrollTop: targetPosition - 50 }, 'slow' );
    });

    const nav = $('#mainNav');
    const navTop = nav.offset().top;
    
    $(window).on("scroll", stickyNavigation);

    function stickyNavigation() {
    	var body = $("body");

		if($(window).scrollTop() >= navTop) {
			body.css("padding-top", nav.outerHeight() + "px");
			nav.addClass("fixed-top");
		}
		else {
			body.css("padding-top", 0);
			nav.removeClass("fixed-top");
		}
    }
});