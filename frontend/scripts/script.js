const recipeList = document.getElementById('recipeList');
const addBtn = document.getElementById('addButton');
const closeBtn = document.querySelector('.close-btn'); // або getElementById('close-btn')
const addForm = document.getElementById('addRecipForm');
const titleInput = document.getElementById('titleInput');
const descriptionInput = document.getElementById('descriptionInput');

window.addEventListener('DOMContentLoaded', () => {
    fetchRecipes();
});

async function fetchRecipes() {
    try {
        const response = await fetch('http://localhost:8080/cook');

        if (!response.ok) {
            throw new Error(`Error: ${response.status}`);
        }

        const recipes = await response.json();
        renderRecipes(recipes);
    } catch (error) {
        console.error('Error loading recipes: ', error);
    }
}

function renderRecipes(recipes) {
    recipeList.innerHTML = '';

    recipes.forEach(recipe => {
        const li = document.createElement('li');
        li.textContent = `${recipe.title} --- ${recipe.description}`;
        recipeList.appendChild(li);
    });
}

addForm.addEventListener('submit', async (e) => {
    e.preventDefault();

    const title = titleInput.value.trim();
    const description = descriptionInput.value.trim();

    try {
        const response = await fetch('http://localhost:8080/cook', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ title, description })
        });

        if (!response.ok) throw new Error(`Error: ${response.status}`);

        titleInput.value = '';
        descriptionInput.value = '';

        fetchRecipes();

    } catch (error) {
        console.error("Error creating recipe:", error);
        alert("Failed creating recipe");
    }
});

addBtn.addEventListener('click', () => {
    addForm.classList.remove('hidden');
});

closeBtn.addEventListener('click', () => {
    addForm.classList.add('hidden');
});