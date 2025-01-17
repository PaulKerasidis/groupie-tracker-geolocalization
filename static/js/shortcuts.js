const cards = document.querySelectorAll('.card');
let selectedCardIndex = 0;
let selectedSugIndex = -1;

document.addEventListener('keydown', (event) => {
    if (event.key === 'ArrowRight') {
        selectedCardIndex = (selectedCardIndex + 1) % cards.length;
    } else if (event.key === 'ArrowLeft') {
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

// Shortcuts for suggestions
document.addEventListener('keydown', (event) => {
    const suggestions = document.querySelectorAll('.suggestion-item');
    if (suggestions.length === 0) return;

    if (event.key === 'ArrowDown') {
        selectedSugIndex = (selectedSugIndex + 1) % suggestions.length;
    } else if (event.key === 'ArrowUp') {
        selectedSugIndex = (selectedSugIndex - 1 + suggestions.length) % suggestions.length;
    }

    suggestions.forEach((suggestion, index) => {
        suggestion.classList.toggle('selected', index === selectedSugIndex);
    });

    // Scroll the selected suggestion into view
    if (selectedSugIndex >= 0) {
        suggestions[selectedSugIndex].scrollIntoView({
            behavior: 'smooth',
            block: 'nearest'
        });
    }
});

document.addEventListener('keydown', (event) => {
    if (event.key === 'Enter' && event.shiftKey && cards[selectedCardIndex]) {
        const idStr = cards[selectedCardIndex].getAttribute('onclick');
        const id = idStr.match(/(\d+)/)[0];
        const artistID = id;
        openPopup(Number(artistID));
    } else if (event.key === 'Enter' && selectedSugIndex >= 0) {
        const suggestions = document.querySelectorAll('.suggestion-item');
        if (suggestions[selectedSugIndex]) {
            suggestions[selectedSugIndex].click();
        }
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