const HOST = "https://faas-nyc1-2ef2e6cc.doserverless.co";
const BASE_PATH = "api/v1/web/fn-cf616548-9087-4219-a8d4-07173ff27fb0/possohelp/api";
const DOWNLOAD_BASE_PATH = `${HOST}/${BASE_PATH}/data/download`;

document.addEventListener('DOMContentLoaded', function() {

  const input           = document.getElementById('phone-number-input');
  const find_button     = document.getElementById('find-button');
  const download_button = document.getElementById('download-button');
  const download_links  = document.getElementById('download-links');

  const dl_births      = document.getElementById("download-link-births");
  const dl_deaths      = document.getElementById("download-link-deaths");
  const dl_rain        = document.getElementById("download-link-rain");
  const dl_temperature = document.getElementById("download-link-temperature");


  input.addEventListener('keyup', function(event) {
    if (event.key === 'Enter') {
      Find(input.value);
    }
    dl_births.href      = `${DOWNLOAD_BASE_PATH}/births/${input.value}`;
    dl_deaths.href      = `${DOWNLOAD_BASE_PATH}/deaths/${input.value}`;
    dl_rain.href        = `${DOWNLOAD_BASE_PATH}/rain/${input.value}`;
    dl_temperature.href = `${DOWNLOAD_BASE_PATH}/temperature/${input.value}`;
  });

  find_button.addEventListener('click', function() {
    Find(input.value);
  });

  download_button.addEventListener('click', function() {
    toggleDisplay(download_links);
  });
  
});

function toggleDisplay(element) {
  if (element.style.display === "none") {
    // Or "flex", "grid", etc. depending on your layout
    element.style.display = "block"; 
  } else {
    element.style.display = "none";
  }
}

function Find(phone_number) {
  Get("", "rain",        phone_number, ChartRain);
  Get("", "temperature", phone_number, ChartTemperature);
  Get("", "births",      phone_number, ChartBirths);
  Get("", "deaths",      phone_number, ChartDeaths);
}

function GetValuesByMonth(rawData, field) {
  var xValues = [];
  var yValues = [];
  for(let i=0; i < rawData.length; i++) {
    const yearMonth = rawData[i].date.slice(0, 7);
    let index = xValues.indexOf(yearMonth);
    if (index < 0) {
      xValues.push(yearMonth);
      yValues.push(rawData[i][field]);
    } else {
      yValues[index] += rawData[i][field];
    }
  }
  return {xValues:xValues, yValues:yValues}
}
