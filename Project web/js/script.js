// Catalog
var catItem = document.getElementsByClassName('catalog');
catItem[0].onclick = function() {
    if (catItem[0].className == 'catalog close') {
        catItem[0].className = 'catalog open';
    }
    else {
        catItem[0].className = 'catalog close';
    }
}

