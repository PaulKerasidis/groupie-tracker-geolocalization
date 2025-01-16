const API_KEY = '0e2H4ClkzL4wy9qJiy6l';

// Get locations from URL parameters
const urlParams = new URLSearchParams(window.location.search);
const locationsParam = urlParams.get('locations');
const cities = locationsParam ? JSON.parse(decodeURIComponent(locationsParam)) : [];

// Initialize map
const map = new maplibregl.Map({
    container: 'map',
    style: `https://api.maptiler.com/maps/streets/style.json?key=${API_KEY}`,
    center: [0, 0],
    zoom: 2
});

// Function to get coordinates for a city
async function getCityCoordinates(cityName) {
    const url = `https://api.maptiler.com/geocoding/${encodeURIComponent(cityName)}.json?key=${API_KEY}`;
    const response = await fetch(url);
    const data = await response.json();
    
    if (data.features && data.features.length > 0) {
        return data.features[0].center;
    }
    return null;
}

// Function to add marker for a city
function addCityMarker(coordinates, cityName) {
    new maplibregl.Marker({
        color: "#FF0000",
        draggable: true
    })
    .setLngLat(coordinates)
    .setPopup(new maplibregl.Popup()
        .setHTML(`
            <h3>${cityName}</h3>
            <p>Coordinates: ${coordinates[0].toFixed(4)}, ${coordinates[1].toFixed(4)}</p>
        `))
    .addTo(map);
}

// Function to calculate bounds for multiple points
function calculateBounds(coordinates) {
    if (!coordinates || coordinates.length === 0) return null;

    const bounds = new maplibregl.LngLatBounds();
    coordinates.forEach(coord => {
        bounds.extend(coord);
    });
    return bounds;
}

// Function to display all cities on the map
async function displayCities(cities) {
    const coordinates = [];
    
    for (const city of cities) {
        const coords = await getCityCoordinates(city);
        if (coords) {
            coordinates.push(coords);
            addCityMarker(coords, city);
        }
    }

    if (coordinates.length > 0) {
        const bounds = calculateBounds(coordinates);
        if (bounds) {
            map.fitBounds(bounds, {
                padding: 50,
                maxZoom: 15
            });
        }
    }
}

// Function to get user location
function getUserLocation() {
    if (!navigator.geolocation) {
        console.error('Geolocation is not supported');
        return;
    }

    navigator.geolocation.getCurrentPosition(
        (position) => {
            const userCoords = [position.coords.longitude, position.coords.latitude];
            addCityMarker(userCoords, 'Your Location');
            map.flyTo({
                center: userCoords,
                zoom: 12
            });
        },
        (error) => {
            console.error('Error getting location:', error.message);
        }
    );
}

// Add map controls
map.addControl(new maplibregl.NavigationControl()); // Navigation
map.addControl(new maplibregl.FullscreenControl()); // Fullscreen
map.addControl(new maplibregl.ScaleControl()); // Scale

// Add user location control
map.addControl(new maplibregl.GeolocateControl({
    positionOptions: {
        enableHighAccuracy: true
    },
    trackUserLocation: true,
    showUserHeading: true
}));

// Add click event listener
map.on('click', (e) => {
    console.log(`Clicked at: ${e.lngLat.lng}, ${e.lngLat.lat}`);
});

// Display cities from URL parameters
displayCities(cities);

// Add map load event
map.on('load', () => {
    console.log('Map loaded successfully');
});

// Add error handling
map.on('error', (e) => {
    console.error('Map error:', e.error);
});