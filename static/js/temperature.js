function ChartTemperature(rawData) {
  console.log("ChartTemperatur", rawData);
  const values = GetValuesByMonth(rawData, "temperature");
  new Chart("chart-temperature", {
    type: "line",
    data: {
      labels: values.xValues,
      datasets: [{
        backgroundColor:"#B98412",
        borderColor: "#96640F",
        data: values.yValues
      }]
    },
    options: {
      legend: {display: false},
      title: {
        display: true,
        text: "Temperature by Month for 2025"
      },
      scales: {
        yAxes: [{ 
          scaleLabel: {
            display: true,
            labelString: 'Celcius' 
          },
          ticks: {
            beginAtZero: true
          }
        }]
      }
    }
  });
}
