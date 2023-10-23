import React from 'react';
import { Card } from '../Card';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';

import { Bar } from 'react-chartjs-2';

ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
);

export const options = {
  plugins: {
    title: {
      display: true,
      text: 'Balance',
      align: 'start',
      color: '#FFFFFF',
      font: {
        size: 24,
      },
    },
    legend:{
      position:'top',
      align:'end',
      labels:{
        boxWidth: 10,
        boxHeight: 10,
        usePointStyle: true,
        pointStyle: 'circle',
        color: '#FFFFFF',
        font:{
          size: 20,
        },
      },
    },
    tooltip:{
      titleFont:{
        size: 20,
      },
      bodyFont:{
        size: 20,
      },
    }
  },
  responsive: true,
  maintainAspectRatio: false,
  scales: {
    x: {
      stacked: true,
      ticks: {
        color: '#FFFFFF',
        font: {
          size: 20,
        },
      }
    },
    y: {
      stacked: true,
      ticks: {
        color: '#FFFFFF',
        font: {
          size: 20,
        },
      }
    },
  },
};

const labels = ['Asset'];

export const  BalanceChart = ({
  equity,
  debt
}: BalanceInterface) => {
  const data = {
    labels,
    datasets: [
      {
        label: 'Asset',
        borderRadius: 4,
        data: [equity + debt],
        backgroundColor: '#4DB9CF',
        stack: 'Stack 0',
      },
      {
        label: 'Equity',
        borderRadius: 4,
        data: [equity],
        backgroundColor: '#83A128',
        stack: 'Stack 1',
      },
      {
        label: 'Debt',
        borderRadius: 4,
        data: [debt],
        backgroundColor: '#F77F00',
        stack: 'Stack 1',
      },
    ],
  };
  return (
    <Card>
      <Bar 
          options={options} 
          data={data} 
          height={400}
          width={800}
      />
    </Card>
  );
}
