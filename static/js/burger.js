(function(){
  // Attendre que le DOM soit complètement chargé
  document.addEventListener('DOMContentLoaded', function(){
    // Récupérer les références des éléments du menu
    const toggle = document.getElementById('menu-toggle');       // Bouton burger
    const menu = document.getElementById('mobile-menu');         // Panneau du menu mobile
    const overlay = document.getElementById('mobile-menu-overlay'); // Fond sombre
    const close = document.getElementById('menu-close');         // Bouton fermer

    // Ouvrir le menu mobile
    function openMenu(){
      menu.classList.remove('hidden', 'translate-x-full');      // Afficher le menu et le glisser dans la viewport
      overlay.classList.remove('hidden');                       // Afficher l'overlay
      toggle.setAttribute('aria-expanded', 'true');             // Mettre à jour l'état d'accessibilité
      document.body.style.overflow = 'hidden';                  // Bloquer le scroll
    }

    // Fermer le menu mobile
    function closeMenu(){
      menu.classList.add('hidden', 'translate-x-full');        // Masquer et glisser vers la droite
      overlay.classList.add('hidden');                         // Masquer l'overlay
      toggle.setAttribute('aria-expanded', 'false');           // Mettre à jour l'état d'accessibilité
      document.body.style.overflow = '';                        // Débloquer le scroll
    }

    // Attacher les événements aux boutons
    toggle?.addEventListener('click', openMenu);               // Clic sur le burger = ouvrir
    close?.addEventListener('click', closeMenu);               // Clic sur la croix = fermer
    overlay?.addEventListener('click', closeMenu);             // Clic sur l'overlay = fermer

    // Fermer le menu en appuyant sur Échap (accessibilité clavier)
    document.addEventListener('keydown', e => {
      if(e.key === 'Escape' && !menu.classList.contains('hidden')) closeMenu();
    });
  });
})();


