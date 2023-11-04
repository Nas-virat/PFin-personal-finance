
import React from 'react';
import {Bar } from 'react-chartjs-2';
import { Chart, registerables} from 'chart.js';

Chart.register(...registerables);

interface BarChartProps {
    title: string;
    xlabel: string;
    ylabel: string;
    data: number[];
    labels: string[];
    backgroundColor: string;
}



export const BarChart = ({title,xlabel, ylabel, data, labels, backgroundColor }: BarChartProps) => {
    const chartData = {
        labels: labels,
        datasets: [
            {
                label: 'Revenue',
                data: data,
                backgroundColor: backgroundColor,
                borderColor: backgroundColor,
                borderWidth: 1,
            },
        ],
    };

    const options = {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
            x: {
                beginAtZero: true,
                title: {
                    display: true,
                    text: xlabel,
                    font: {
                        size: 20, // Set x-axis label size
                    },
                    color: '#EAE9E9', // Change x-axis label color
                },
                ticks: {
                    font: {
                        size: 20, // Set x-axis ticks size
                    },
                    color: '#EAE9E9', // Change x-axis ticks color
                },

            },
            y: {
                beginAtZero: true,
                title: {
                    display: true,
                    text: ylabel,
                    font: {
                        size: 20, // Set y-axis label size
                    },
                    color: '#EAE9E9', // Change y-axis label color
                },
                ticks: {
                    font: {
                        size: 20, // Set x-axis ticks size
                    },
                    color: '#EAE9E9', // Change x-axis ticks color
                },
            },
        },
        plugins: {
            title: {
                display: true,
                text: title,
                align: 'start',
                color: '#EAE9E9',
                font: {
                  size: 24,
                },
              },
            legend: {
                labels: {
                    boxWidth: 10,
                    boxHeight: 10,
                    usePointStyle: true,
                    pointStyle: 'circle',
                    font: {
                        size: 20, // Change legend label size
                    },
                    color: '#EAE9E9', // Change legend label color
                },
            },
            tooltip:{
                titleFont:{
                  size: 20,
                },
                bodyFont:{
                  size: 20,
                },
            },
        },
    };

    return (
        <div className='w-full flex justify-center'>
            <Bar data={chartData} options={options} height={400} />
        </div>
    );
};