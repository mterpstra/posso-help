function ChartBirths(rawData) {
  const values = aggregateBirthData(rawData);
  var barColors = ["#92C5F9","#4394E5","#0066CC","#004D99","#003366"];

  new Chart("chart-births", {
    type: "pie",
    data: {
      labels: values.xValues,
      datasets: [{
        backgroundColor: barColors,
        data: values.yValues
      }]
    },
    options: {
      title: {
        display: true,
        text: "Births by Breed"
      }
    }
  });
}

function aggregateBirthData(rawData) {
  var xValues = [];
  var yValues = [];
  for(let i=0; i < rawData.length; i++) {
    let index = xValues.indexOf(rawData[i].breed);
    if (index < 0) {
      xValues.push(rawData[i].breed);
      yValues.push(1);
    } else {
      yValues[index]++;
    }
  }
  return {xValues:xValues, yValues:yValues}
}

/*
function ChartBirths(rawData) {
  console.log("ChartBirths", rawData);
  const xValues = [100,200,300,400,500,600,700,800,900,1000];

  new Chart("chart-births", {
    type: "line",
    data: {
      labels: xValues,
      datasets: [{
        data: [860,1140,1060,1060,1070,1110,1330,2210,7830,2478],
        borderColor: "red",
        fill: false
      },{
        data: [1600,1700,1700,1900,2000,2700,4000,5000,6000,7000],
        borderColor: "green",
        fill: false
      },{
        data: [300,700,2000,5000,6000,4000,2000,1000,200,100],
        borderColor: "blue",
        fill: false
      }]
    },
    options: {
      legend: {display: false},
      title: {
        display: true,
        text: "Births by Month for 2025"
      }
    }
  });
}
*/
