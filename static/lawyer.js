document.addEventListener('DOMContentLoaded', function() {
    // Fetch and display lawyer profiles
    fetch('/lawyer/')
        .then(response => response.json())
        .then(data => {
            const lawyerList = document.getElementById('lawyers');
            data.forEach(lawyer => {
                const div = document.createElement('div');
                div.innerHTML = `
                    <h2>${lawyer.name}</h2>
                    <p>Specialty: ${lawyer.specialty}</p>
                    <p>Contact: ${lawyer.contact_info}</p>
                `;
                lawyerList.appendChild(div);
            });
        })
        .catch(error => console.error('Error fetching lawyer profiles:', error));
});
