function ChartDeaths(rawData) {
  console.log("ChartDeaths", rawData);
  const values = aggregateDeathData(rawData);
  var barColors = ["#F89B78","#F4784A","#F0561D","#B1380B","#731F00"];

  new Chart("chart-deaths", {
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
        text: "Deaths by Cause"
      }
    }
  });
}

function aggregateDeathData(rawData) {
  var xValues = [];
  var yValues = [];
  for(let i=0; i < rawData.length; i++) {
    let index = xValues.indexOf(rawData[i].cause);
    if (index < 0) {
      xValues.push(rawData[i].cause);
      yValues.push(1);
    } else {
      yValues[index]++;
    }
  }
  return {xValues:xValues, yValues:yValues}
}
