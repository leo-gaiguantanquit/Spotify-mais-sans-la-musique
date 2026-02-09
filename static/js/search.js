// Attend que le DOM soit entièrement chargé avant d'exécuter le script
document.addEventListener('DOMContentLoaded', function() {
    const searchInput = document.getElementById('search-input');
    
    // Création du conteneur pour les résultats de recherche
    const resultsContainer = document.createElement('div');
    resultsContainer.id = 'search-results';
    resultsContainer.className = 'search-results-dropdown';
    
    // Insertion du conteneur après le champ de saisie
    searchInput.parentNode.appendChild(resultsContainer);
    
    // Application des styles pour le conteneur des résultats
    // Ces styles positionnent la liste déroulante sous la barre de recherche
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

    // Écouteur d'événement sur la saisie dans le champ de recherche
    searchInput.addEventListener('input', function(e) {
        const query = e.target.value;
        
        // Annule le timer précédent pour éviter les requêtes multiples (Debounce)
        clearTimeout(debounceTimer);
        
        // Si la requête est trop courte, on cache les résultats
        if (query.length < 2) {
            resultsContainer.style.display = 'none';
            return;
        }

        // Attend 300ms après la fin de la saisie avant d'envoyer la requête
        debounceTimer = setTimeout(() => {
            fetch(`/api/search?q=${encodeURIComponent(query)}`)
                .then(response => response.json())
                .then(data => {
                    displayResults(data);
                })
                .catch(err => {
                    console.error('Erreur lors de la récupération des résultats :', err);
                });
        }, 300);
    });

    // Ferme les résultats si on clique en dehors du champ ou de la liste
    document.addEventListener('click', function(e) {
        if (e.target !== searchInput && e.target !== resultsContainer && !resultsContainer.contains(e.target)) {
            resultsContainer.style.display = 'none';
        }
    });

    // Fonction pour afficher les résultats de la recherche dans le DOM
    function displayResults(data) {
        resultsContainer.innerHTML = '';
        
        // Si aucune donnée ou aucun artiste trouvé, on cache le conteneur
        if (!data || !data.artists || !data.artists.items || data.artists.items.length === 0) {
            resultsContainer.style.display = 'none';
            return;
        }

        resultsContainer.style.display = 'block';

        const list = document.createElement('ul');
        list.style.listStyle = 'none';
        list.style.margin = '0';
        list.style.padding = '0';

        // Limite l'affichage aux 8 premiers résultats
        const artists = data.artists.items.slice(0, 8);

        artists.forEach(artist => {
            const item = document.createElement('li');
            item.className = 'search-result-item';
            
            // Création du lien vers la page de l'artiste
            const link = document.createElement('a');
            link.href = `/artiste/${artist.id}`;
            link.style.display = 'flex';
            link.style.alignItems = 'center';
            link.style.padding = '10px';
            link.style.textDecoration = 'none';
            link.style.color = '#fff';
            link.style.borderBottom = '1px solid #333';
            link.style.transition = 'background-color 0.2s';

            // Effet de survol
            link.onmouseover = function() { this.style.backgroundColor = '#333'; };
            link.onmouseout = function() { this.style.backgroundColor = 'transparent'; };

            // Gestion de l'image de l'artiste
            const img = document.createElement('img');
            if (artist.images && artist.images.length > 0) {
                // Utilise la plus petite image disponible
                img.src = artist.images[artist.images.length - 1].url;
            } else {
                img.src = 'https://placehold.co/40x40';
            }
            img.style.width = '40px';
            img.style.height = '40px';
            img.style.borderRadius = '50%';
            img.style.marginRight = '10px';
            img.style.objectFit = 'cover';

            // Conteneur pour le nom et le genre
            const infoDiv = document.createElement('div');
            infoDiv.style.display = 'flex';
            infoDiv.style.flexDirection = 'column';

            const name = document.createElement('span');
            name.textContent = artist.name;
            name.style.fontWeight = 'bold';
            name.style.fontSize = '14px';

            const genre = document.createElement('span');
            genre.textContent = "Artiste"; 
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
