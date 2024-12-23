document.addEventListener("DOMContentLoaded", () => {
    const searchInput = document.getElementById("q");
    const suggestionsBox = document.getElementById("suggestions");
  
    searchInput.addEventListener("input", async () => {
      const query = searchInput.value.trim();
      if (query.length < 1) {
        suggestionsBox.innerHTML = "";
        return;
      }
  
      try {
        const response = await fetch(`/suggestions?q=${encodeURIComponent(query)}`);
        if (!response.ok) {
          throw new Error("Failed to fetch suggestions");
        }
  
        const suggestions = await response.json();
        displaySuggestions(suggestions);
      } catch (error) {
        console.error(error);
        suggestionsBox.innerHTML = "<div class='suggestion-item'>Error loading suggestions</div>";
      }
    });
  
    function displaySuggestions(suggestions) {
      if (suggestions.length === 0) {
        suggestionsBox.innerHTML = "<div class='suggestion-item'>No results found</div>";
        return;
      }
  
      suggestionsBox.innerHTML = suggestions
        .slice(0, 20) // Add this line to limit to 10 suggestions
        .map(item => `<div class='suggestion-item' data-id='${item.id}'>${item.value} - ${item.type}</div>`)
        .join("");
  
      document.querySelectorAll(".suggestion-item").forEach(item => {
        item.addEventListener("click", () => {
          searchInput.value = item.textContent.split(" - ")[0];
          suggestionsBox.innerHTML = "";
        });
      });
    }
  });

  
