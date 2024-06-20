document.addEventListener("DOMContentLoaded", () => {
    const nameContainer = document.getElementById('emoji-container');
    nameContainer.style.zIndex = -1
    const names = [
        'A. Nassuif', 'Vivien', 'Anne-Marie', 'Adrien', 'Ajoly', 'Alexandre',
        'Antoine M', 'Arnaud', 'Arrachid', 'Nattan', 'Pierre', 'Nicolas', 'Nathan',
        'Romain', 'Sam', 'Theo', 'Thomas', 'Tichane', 'Valeryan', 'Valentin', 'Clement',
        'Daro', 'Djihadi', 'Dylan B.', 'Jaouhar76', 'Lucas',
        'Mohammed', 'mouldmoh', 'A. Martini', 'Adam', 'Adrien', 'Bastien', 'Leo', 'Bob',
        'Brendie', 'Christophe', 'Damien', 'Djihadi', 'Esteban', 'Fabien', 'Fares', 'Gabriel',
        'JLasseny', 'Justine', 'Maxime', 'Nicolas', 'Toni', 'Woodkens', 'Yanis', 'Yoann'
      ];

    function createName() {
        const name = document.createElement('div');
        name.classList.add('emoji');
        name.style.color = getRandomColor();  // Assign random color to emoji
        name.textContent = names[Math.floor(Math.random() * names.length)];
        name.style.left = Math.random() * 90 + 'vw';
        name.style.animationDuration = Math.random() * 5 + 7 + 's';  // Random duration between 7s and 12s
        nameContainer.appendChild(name);

        // Remove emoji after animation ends to prevent clutter
        name.addEventListener('animationend', () => {
            name.remove();
        });
    }

    function getRandomColor() {
        const colors = ['#FF0000', '#00FF00', '#0000FF', '#FF00FF', '#FFFF00', '#00FFFF'];
        return colors[Math.floor(Math.random() * colors.length)];
    }


    // Create emojis at intervals
    setInterval(createName, 1000);  // Change interval as needed
});
