import React from 'react';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Doughnut } from 'react-chartjs-2';

ChartJS.register(ArcElement, Tooltip, Legend);


const legendOptions = {
  display: true,
  position: 'right', // You can change the position if needed
  labels: {
    font: {
      size: 20, // Change the font size here
    },
    color: '#EAE9E9', // Change the font color here
  },
};


const options = {
  plugins: {
    legend: legendOptions,
  },
};


 // Define the CSS styles for the chart container
 const chartContainerStyle = {
  width: '320px', // Change the width as desired
  height: '320px', // Change the height as desired
};


export const DoughnutChart = ({data,labels}: DoughnutChartInterface) => {
  
  const chartdata = {
    labels: labels,
    datasets: [
      {
        label: 'Amount',
        data: data,
        backgroundColor: [
          'rgba(255, 99, 132, 1)',
          'rgba(54, 162, 235, 1)',
          'rgba(255, 206, 86, 1)',
          'rgba(75, 192, 192, 1)',
          'rgba(153, 102, 255, 1)',
          'rgba(255, 159, 64, 1)',
        ],
        borderColor: [
          'rgba(255, 99, 132, 1)',
          'rgba(54, 162, 235, 1)',
          'rgba(255, 206, 86, 1)',
          'rgba(75, 192, 192, 1)',
          'rgba(153, 102, 255, 1)',
          'rgba(255, 159, 64, 1)',
        ],
        borderWidth: 1,
      },
    ],
  }
  return (
    <div style={chartContainerStyle}>
      <Doughnut data={chartdata} options={options} />
    </div>

  );
}
