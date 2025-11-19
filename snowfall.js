document.addEventListener("DOMContentLoaded", function() {
    const svg = document.getElementById('tree');

    // Function to create snowflakes
    function createSnowflakes(count) {
        for (let i = 0; i < count; i++) {
            // Define snowflake properties
            const snowflake = document.createElementNS("http://www.w3.org/2000/svg", "circle");
            const snowflakeX = Math.random() * window.innerWidth; // Random x position across the window width
            const snowflakeY = Math.random() * -100; // Start above the viewport
            const snowflakeSize = Math.random() * 8 + 2; // Random size between 2 and 10

            // Set snowflake attributes
            snowflake.setAttribute("cx", snowflakeX);
            snowflake.setAttribute("cy", snowflakeY);
            snowflake.setAttribute("r", snowflakeSize);
            snowflake.setAttribute("fill", "white");

            // Append snowflake to the SVG
            svg.appendChild(snowflake);

            // Animate the snowflake falling
            animateSnowflake(snowflake);
        }
    }

    // Function to animate a snowflake falling
    function animateSnowflake(snowflake) {
        let snowflakeY = parseFloat(snowflake.getAttribute("cy"));
        const fallSpeed = Math.random() * 2 + 1; // Random fall speed

        function fall() {
            snowflakeY += fallSpeed; // Update the y position
            if (snowflakeY > window.innerHeight) { // Reset to the top if it falls below the viewport height
                snowflakeY = -10; // Start above the viewport
                const newX = Math.random() * window.innerWidth; // Random new x position across the window
                snowflake.setAttribute("cx", newX); // Update x position
            }
            snowflake.setAttribute("cy", snowflakeY); // Update y position
            requestAnimationFrame(fall); // Continue the animation
        }

        fall(); // Start the falling animation
    }

    // Create snowflakes
    createSnowflakes(50); // Create 50 snowflakes
});
