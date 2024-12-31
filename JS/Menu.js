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

// 選取元素
const menuItems = document.querySelectorAll(".menu-item");
const popup = document.getElementById("popup");
const closePopupButton = document.getElementById("closePopup");
const popupTitle = document.getElementById("popupTitle");
const popupDescription = document.getElementById("popupDescription");
const popupPrice = document.getElementById("popupPrice");
const confirmOrder = document.getElementById("confirmOrder");

// 點擊餐點顯示彈窗
menuItems.forEach((item) => {
  item.addEventListener("click", () => {
    const name = item.getAttribute("data-name");
    const price = item.getAttribute("data-price");
    const description = item.getAttribute("data-desc");

    // 更新彈窗內容
    popupTitle.textContent = name;
    popupDescription.textContent = description;
    popupPrice.textContent = price;

    // 顯示彈窗
    popup.style.display = "flex";
  });
});

// 關閉彈窗
closePopupButton.addEventListener("click", () => {
  popup.style.display = "none";
});

// 點擊外部關閉彈窗
window.addEventListener("click", (event) => {
  if (event.target === popup) {
    popup.style.display = "none";
  }
});

// 確認訂單按鈕
confirmOrder.addEventListener("click", () => {
  alert("已加入購物車！");
  popup.style.display = "none";
});
