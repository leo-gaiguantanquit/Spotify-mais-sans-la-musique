document.addEventListener('DOMContentLoaded', function() {
    const searchInput = document.getElementById('search-input');
    const resultsContainer = document.createElement('div');
    resultsContainer.id = 'search-results';
    resultsContainer.className = 'search-results-dropdown';
    
    // Insert results container after the input
    searchInput.parentNode.appendChild(resultsContainer);
    
    // Style for the results container (you can move this to style.css)
    resultsContainer.style.position = 'absolute';
    resultsContainer.style.width = '300px';
    resultsContainer.style.maxHeight = '400px';
    resultsContainer.style.overflowY = 'auto';
    resultsContainer.style.backgroundColor = '#282828';
    resultsContainer.style.borderRadius = '0 0 8px 8px';
    resultsContainer.style.zIndex = '1000';
    resultsContainer.style.display = 'none';
    resultsContainer.style.boxShadow = '0 4px 6px rgba(0,0,0,0.3)';
    resultsContainer.style.top = '100%';
    resultsContainer.style.left = '0';

    let debounceTimer;

    searchInput.addEventListener('input', function(e) {
        const query = e.target.value;
        
        clearTimeout(debounceTimer);
        
        if (query.length < 2) {
            resultsContainer.style.display = 'none';
            return;
        }

        debounceTimer = setTimeout(() => {
            fetch(`/api/search?q=${encodeURIComponent(query)}`)
                .then(response => response.json())
                .then(data => {
                    displayResults(data);
                })
                .catch(err => {
                    console.error('Error fetching search results:', err);
                });
        }, 300); // 300ms debounce
    });

    // Close results when clicking outside
    document.addEventListener('click', function(e) {
        if (e.target !== searchInput && e.target !== resultsContainer && !resultsContainer.contains(e.target)) {
            resultsContainer.style.display = 'none';
        }
    });

    function displayResults(data) {
        resultsContainer.innerHTML = '';
        
        if (!data || !data.artists || !data.artists.items || data.artists.items.length === 0) {
            resultsContainer.style.display = 'none';
            return;
        }

        resultsContainer.style.display = 'block';

        const list = document.createElement('ul');
        list.style.listStyle = 'none';
        list.style.margin = '0';
        list.style.padding = '0';

        // Limit to 5 results
        const artists = data.artists.items.slice(0, 8);

        artists.forEach(artist => {
            const item = document.createElement('li');
            item.className = 'search-result-item';
            
            const link = document.createElement('a');
            link.href = `/artiste/${artist.id}`; // Adjusted to match your route structure
            link.style.display = 'flex';
            link.style.alignItems = 'center';
            link.style.padding = '10px';
            link.style.textDecoration = 'none';
            link.style.color = '#fff';
            link.style.borderBottom = '1px solid #333';
            link.style.transition = 'background-color 0.2s';

            link.onmouseover = function() { this.style.backgroundColor = '#333'; };
            link.onmouseout = function() { this.style.backgroundColor = 'transparent'; };

            // Image
            const img = document.createElement('img');
            if (artist.images && artist.images.length > 0) {
                img.src = artist.images[artist.images.length - 1].url; // Smallest image
            } else {
                img.src = 'https://placehold.co/40x40';
            }
            img.style.width = '40px';
            img.style.height = '40px';
            img.style.borderRadius = '50%';
            img.style.marginRight = '10px';
            img.style.objectFit = 'cover';

            // Name and Info
            const infoDiv = document.createElement('div');
            infoDiv.style.display = 'flex';
            infoDiv.style.flexDirection = 'column';

            const name = document.createElement('span');
            name.textContent = artist.name;
            name.style.fontWeight = 'bold';
            name.style.fontSize = '14px';

            const genre = document.createElement('span');
            genre.textContent = "Artiste"; // You can show genre if available
            genre.style.fontSize = '12px';
            genre.style.color = '#b3b3b3';

            infoDiv.appendChild(name);
            infoDiv.appendChild(genre);

            link.appendChild(img);
            link.appendChild(infoDiv);
            item.appendChild(link);
            list.appendChild(item);
        });
        
        resultsContainer.appendChild(list);
    }
});
