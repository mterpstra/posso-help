function ChartRain(rawData) {
  const values = GetValuesByMonth(rawData, "amount");
  var barColors = ["#9AD8D8","#63BDBD","#37A3A3","#147878","#004D4D"];
  new Chart("chart-rain", {
    type: "bar",
    data: {
      labels: values.xValues,
      datasets: [{
        backgroundColor: barColors,
        data: values.yValues
      }]
    },
    options: {
      legend: {display: false},
      title: {
        display: true,
        text: "Rain by Month in 2025"
      },
      scales: {
        yAxes: [{ 
          scaleLabel: {
            display: true,
            labelString: 'Millimeters' 
          },
          ticks: {
            beginAtZero: true
          }
        }]
      }
    }
  });
}
