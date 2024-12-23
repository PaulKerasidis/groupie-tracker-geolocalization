const key = '0e2H4ClkzL4wy9qJiy6l';

// Get locations from URL parameters
const urlParams = new URLSearchParams(window.location.search);
const locationsParam = urlParams.get('locations');
const cities = locationsParam ? JSON.parse(decodeURIComponent(locationsParam)) : [];

// Initialize map
const map = new maplibregl.Map({
    container: 'map',
    style: `https://api.maptiler.com/maps/streets/style.json?key=${key}`,
    center: [0, 0],
    zoom: 2
});

// Function to get coordinates for a city
async function getCityCoordinates(cityName) {
    const url = `https://api.maptiler.com/geocoding/${encodeURIComponent(cityName)}.json?key=${key}`;
    const response = await fetch(url);
    const data = await response.json();
    
    if (data.features && data.features.length > 0) {
        return data.features[0].center;
    }
    return null;
}

// Function to add marker for a city
function addCityMarker(coordinates, cityName) {
    new maplibregl.Marker({color: "#FF0000"})
        .setLngLat(coordinates)
        .setPopup(new maplibregl.Popup().setHTML(`<h3>${cityName}</h3>`))
        .addTo(map);
}

// Display cities from URL parameters
async function displayCities(cities) {
    for (const city of cities) {
        const coords = await getCityCoordinates(city);
        if (coords) {
            addCityMarker(coords, city);
        }
    }
}

// Initialize map with locations
displayCities(cities);
map.addControl(new maplibregl.NavigationControl());