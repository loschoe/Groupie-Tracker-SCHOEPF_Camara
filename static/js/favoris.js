function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
    return '';
}

function setCookie(name, value, days = 365) {
    const expires = new Date(Date.now() + days * 24 * 60 * 60 * 1000).toUTCString();
    document.cookie = `${name}=${value}; expires=${expires}; path=/`;
}

function getFavorites() {
    const favorites = getCookie('favorites');
    return favorites ? favorites.split(',').filter(id => id !== '').map(id => parseInt(id)) : []; 
}

function saveFavorites(favorites) {
    setCookie('favorites', favorites.join(','));
}

function toggleFavorite(artistId) {
    let favorites = getFavorites();
    const index = favorites.indexOf(artistId);

    if (index > -1) {
        favorites.splice(index, 1);
    } else {
        favorites.push(artistId);
    }

    saveFavorites(favorites);
    updateFavoriteButtons();
    
    if (window.location.pathname === '/favoris') {
        setTimeout(() => {
            location.reload();
        }, 300);
    }
}

function updateFavoriteButtons() {
    const favorites = getFavorites();
    const buttons = document.querySelectorAll('.favorite-btn');

    buttons.forEach(btn => {
        const artistId = parseInt(btn.getAttribute('data-id'));
        if (favorites.includes(artistId)) {
            btn.classList.add('active');
            btn.innerHTML = '<i class="fa-solid fa-heart"></i>';
        } else {
            btn.classList.remove('active');
            btn.innerHTML = '<i class="fa-regular fa-heart"></i>';
        }
    });

    updateFavoriteCount();
}

function updateFavoriteCount() {
    const count = getFavorites().length;
    const navLinks = document.querySelectorAll('nav a[href="/favoris"]');
    
    navLinks.forEach(navLink => {
        const span = navLink.querySelector('span');
        if (span) {
            if (count > 0) {
                span.textContent = `Favoris (${count})`;
            } else {
                span.textContent = 'Favoris';
            }
        }
    });
}

document.addEventListener('DOMContentLoaded', () => {
    updateFavoriteButtons();
});
