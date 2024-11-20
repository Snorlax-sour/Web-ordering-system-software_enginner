// script.js
function showCategory(category) {
    // 移除所有菜單項目
    const menuItems = document.querySelectorAll('.menu-item');
    menuItems.forEach(item => {
        if (item.getAttribute('data-category') === category) {
            item.style.display = 'flex'; // 顯示對應分類
        } else {
            item.style.display = 'none'; // 隱藏其他分類
        }
    });

    // 更新分類標籤的樣式
    const categories = document.querySelectorAll('.category');
    categories.forEach(cat => {
        if (cat.textContent === category) {
            cat.classList.add('active');
        } else {
            cat.classList.remove('active');
        }
    });
}