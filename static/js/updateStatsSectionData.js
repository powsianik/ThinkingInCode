function updateStatsSectionElements() {
	var commercialExperience = getCommercialExperienceYears();
    var nonCommercialExperience = getNonCommercialExperienceYears();

    var commercialElement = document.getElementById('commercial-experience');
    var nonCommercialElement = document.getElementById('non-commercial-experience');

    if (commercialElement && nonCommercialElement) {

        commercialElement.textContent = commercialExperience;
        commercialElement.setAttribute('data-to', commercialExperience);

        nonCommercialElement.textContent = nonCommercialExperience;
        nonCommercialElement.setAttribute('data-to', nonCommercialExperience);
    }
}