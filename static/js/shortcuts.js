const cards = document.querySelectorAll('.card');
let id = 0;
let selectedCardIndex = 0;

document.addEventListener('keydown', (event) => {
    if (event.key === 'ArrowRight') {
        selectedCardIndex = (selectedCardIndex + 1) % cards.length;
    } else if (event.key === 'ArrowLeft' ) {
        selectedCardIndex = (selectedCardIndex - 1 + cards.length) % cards.length;
    }
    cards.forEach((card, index) => {
        card.classList.toggle('selected', index === selectedCardIndex);
    });
    
    // Scroll the selected card into view
    cards[selectedCardIndex].scrollIntoView({
        behavior: 'smooth',
        block: 'nearest'
    });
});

document.addEventListener('keydown', (event) => {
    if (event.key === 'Enter' && event.shiftKey && cards[selectedCardIndex]) {
        idStr = cards[selectedCardIndex].getAttribute('onclick');
        id = idStr.match(/(\d+)/)[0];
        const artistID = id;
        openPopup(Number(artistID));
    }   
});


document.addEventListener('keydown', (event) => {
    if (event.ctrlKey && event.key === 'f') {
        event.preventDefault();
        document.querySelector('#q').focus();
    } else if (event.key === 'Escape') {
        closePopup();
    }
});

