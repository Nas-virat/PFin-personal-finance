import React from 'react';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Doughnut } from 'react-chartjs-2';

import tableau10Colors from '@/config/color';

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
    tooltip:{
      titleFont:{
        size: 20,
      },
      bodyFont:{
        size: 20,
      },
    }
  },
};


 // Define the CSS styles for the chart container
 const chartContainerStyle = {
  width: '350px', // Change the width as desired
  height: '350px', // Change the height as desired
};


export const DoughnutChart = ({data,labels}: DoughnutChartInterface) => {
  
  const chartdata = {
    labels: labels,
    datasets: [
      {
        label: 'Amount',
        data: data,
        backgroundColor: tableau10Colors,
        borderColor: tableau10Colors,
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
