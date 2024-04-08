document.addEventListener("DOMContentLoaded", function() {

    const productContainers = document.querySelectorAll(".product-container");
    const productTitles = document.querySelectorAll(".product-title");
    const deleteBtns = document.querySelectorAll(".delete-btn");
    const editBtns = document.querySelectorAll(".edit-btn");
    const saveBtns = document.querySelectorAll(".save-btn");
    const descriptions = document.querySelectorAll(".description");
    const editDescriptions = document.querySelectorAll(".edit-description");
    const searchInput = document.getElementById("searchInput");
    const createProductBtn = document.getElementById("createProductBtn");
    const sortButton = document.getElementById("sortButton");


    function filterProducts(searchTerm) {
        productContainers.forEach(container => {
            const productName = container.querySelector(".product-title").innerText.toLowerCase();
            const isVisible = productName.includes(searchTerm);
            container.style.display = isVisible ? "block" : "none";
        });
    }


    searchInput.addEventListener("input", function() {
        const searchTerm = searchInput.value.trim().toLowerCase();
        filterProducts(searchTerm);
    });


    productContainers.forEach((container, index) => {
        const deleteBtn = container.querySelector(".delete-btn");

        deleteBtn.addEventListener("click", function() {
            const productId = container.dataset.productId;

            fetch(`/api/products/${productId}`, {
                method: 'DELETE',
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                window.location.reload();
            })
            .catch(error => {
                console.error('There was a problem with your fetch operation:', error);
            });
        });

        productTitles[index].addEventListener("click", function() {
            const description = descriptions[index];
            description.style.display = (description.style.display === "none" || description.style.display === "") ? "block" : "none";
        });

        editBtns[index].addEventListener("click", function() {
            const description = descriptions[index];
            const editDescription = editDescriptions[index];

            description.style.display = "none";
            editDescription.value = description.querySelector("p").innerText;
            editDescription.style.display = "block";
            editBtns[index].style.display = "none";
            saveBtns[index].style.display = "block";
        });

        saveBtns[index].addEventListener("click", function() {
            const productId = container.dataset.productId;
            const editDescription = editDescriptions[index].value;

            fetch(`/api/products/${productId}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    Description: editDescription
                })
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                window.location.reload();
            })
            .catch(error => {
                console.error('There was a problem with your fetch operation:', error);
            });
        });
    });

    const productsPerPage = 2; // Change this value as per your requirement
    const products = document.querySelectorAll('.product-container');
    const pagination = document.querySelector('.pagination');

    function displayProducts(startIndex, endIndex) {
        products.forEach((product, index) => {
            if (index >= startIndex && index < endIndex) {
                product.style.display = 'block';
            } else {
                product.style.display = 'none';
            }
        });
    }

    function createPaginationButtons(totalPages) {
        pagination.innerHTML = '';

        for (let i = 1; i <= totalPages; i++) {
            const pageItem = document.createElement('li');
            pageItem.classList.add('page-item');

            const pageLink = document.createElement('a');
            pageLink.classList.add('page-link');
            pageLink.href = '#';
            pageLink.textContent = i;

            pageItem.appendChild(pageLink);
            pagination.appendChild(pageItem);

            pageLink.addEventListener('click', function() {
                const pageNumber = parseInt(pageLink.textContent);
                const startIndex = (pageNumber - 1) * productsPerPage;
                const endIndex = startIndex + productsPerPage;
                displayProducts(startIndex, endIndex);
            });
        }
    }

    function initPagination() {
        const totalPages = Math.ceil(products.length / productsPerPage);
        createPaginationButtons(totalPages);
        displayProducts(0, productsPerPage);
    }

    initPagination();

    createProductBtn.addEventListener("click", function() {
        const productName = document.getElementById("productName").value;
        const productPrice = document.getElementById("productPrice").value;
        const productDescription = document.getElementById("productDescription").value;
        const productQuantity = document.getElementById("productQuantity").value;
        const productImage = document.getElementById("productImage").value;

        const newProduct = {
            name: productName,
            price: parseFloat(productPrice),
            description: productDescription,
            quantityInStock: parseInt(productQuantity),
            imagePath: productImage
        };

        fetch("/api/products", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(newProduct)
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Failed to create product');
            }
            window.location.reload();
        })
        .catch(error => {
            console.error('Error creating product:', error);
        });
    });

    sortButton.addEventListener("click", function() {
        const products = Array.from(productContainers);

        products.sort((a, b) => {
            const priceA = parseFloat(a.querySelector(".product-price").innerText.split(" ")[1]);
            const priceB = parseFloat(b.querySelector(".product-price").innerText.split(" ")[1]);
            return priceA - priceB;
        });

        const productsSection = document.querySelector('.products');
        productsSection.innerHTML = "";

        products.forEach(product => {
            productsSection.appendChild(product);
        });
    });
});