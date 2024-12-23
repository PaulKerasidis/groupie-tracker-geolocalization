let popupElement = document.getElementById('popup');
let artistList = [];
let relationList = [];
let locationList = [];
let dateList = [];


function openPopup(artistId) {
    const selectedArtist = artistList.find(artist => artist.id === artistId);

    if (selectedArtist) {
        const popupImage = popupElement.querySelector('img');

        const locationGridElement = document.querySelector('.location-grid');
        const locationContainer = document.querySelector('.location-container');
        const dateContainer = document.querySelector('.date-container');
        locationGridElement.innerHTML = '';
        locationContainer.innerHTML = '';
        dateContainer.innerHTML = '';

        // Populate locations
        for (let location of selectedArtist.Locations) {
            const locationDiv = document.createElement('div');
            locationDiv.classList.add('location-item');

            const locationText = document.createElement('p');
            locationText.textContent = location;

            locationDiv.appendChild(locationText);
            locationContainer.appendChild(locationDiv);
        }

        // Populate dates
        for (let date of selectedArtist.Dates) {
            const dateDiv = document.createElement('div');
            dateDiv.classList.add('date-item');

            const dateText = document.createElement('p');
            dateText.textContent = date;

            dateDiv.appendChild(dateText);
            dateContainer.appendChild(dateDiv);
        }

        // Populate location-dates relations
        for (let [location, dates] of selectedArtist.Relations) {
            const locationRelationDiv = document.createElement('div');
            locationRelationDiv.classList.add('location-relation');

            const locationTitle = document.createElement('h3');
            locationTitle.textContent = location;

            const dateListParagraph = document.createElement('p');
            dateListParagraph.textContent = dates.join(', ');

            locationRelationDiv.appendChild(locationTitle);
            locationRelationDiv.appendChild(dateListParagraph);
            locationGridElement.appendChild(locationRelationDiv);
        }

        // Set artist image and alt text
        popupImage.src = selectedArtist.image;
        popupImage.alt = selectedArtist.name;

        popupElement.classList.add('open-popup');
    } else {
        console.error('Artist not found for ID:', artistId);
    }
}
function showLocationsOnMap() {
    const locationContainer = document.querySelector('.location-container');
    const locations = Array.from(locationContainer.children)
        .map(div => div.textContent.trim())
        .filter(loc => loc); // Remove empty locations
    
    if (locations.length > 0) {
        const locationParam = encodeURIComponent(JSON.stringify(locations));
        window.open(`/map?locations=${locationParam}`, '_blank');
    }
}

function closePopup() {
    popupElement.classList.remove('open-popup');
}

async function fetchLocations() {
    try {
        const response = await fetch('/api/locations');
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        locationList = await response.json();
        console.log(locationList);
    } catch (error) {
        console.error('Error fetching artists:', error);
    }
}

async function fetchDates() {
    try {
        const response = await fetch('/api/dates');
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        dateList = await response.json();
        console.log(dateList);
    } catch (error) {
        console.error('Error fetching artists:', error);
    }
}

async function fetchRelations() {
    try {
        const response = await fetch('/api/relations');
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        relationList = await response.json();
        console.log(relationList);
    } catch (error) {
        console.error('Error fetching artists:', error);
    }
}

async function fetchArtists() {
    try {
        const response = await fetch('/api/artists');
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        artistList = await response.json();

        for (let i = 0; i < artistList.length; i++) {
            artistList[i].Relations = new Map(Object.entries(relationList[i]));
            artistList[i].Locations = locationList[i];
            artistList[i].Dates = dateList[i];
        }

        console.log(artistList);
    } catch (error) {
        console.error('Error fetching artists:', error);
    }
}

function initializeApp() {
    fetchRelations();
    fetchLocations();
    fetchDates();
    fetchArtists();
}

initializeApp();
