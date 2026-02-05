
input.addEventListener('input', async function() {
  const query = input.value;
  if (query.length < 2) return; // attendre 2 lettres
  const res = await fetch('/api/search?q=' + encodeURIComponent(query));
  const suggestions = await res.json();
  // afficher suggestions sous le champ
});