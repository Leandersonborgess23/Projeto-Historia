let currentSlide = 0;
    function showSlide(index) {
        const slides = document.querySelectorAll('.slide');
        slides.forEach((slide, i) => {
            slide.classList.toggle('active', i === index);
        });
        currentSlide = index;
    
        window.scrollTo(0, 0);
    }

    function nextSlide() {
        const slides = document.querySelectorAll('.slide');
        if (currentSlide < slides.length - 1) {
            showSlide(currentSlide + 1);
        }
    }
    function prevSlide() {
        if (currentSlide > 0) {
            showSlide(currentSlide - 1);
         }
    }
        document.addEventListener('DOMContentLoaded', () => {
            showSlide(currentSlide);
    });