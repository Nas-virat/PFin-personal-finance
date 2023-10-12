"use client";
import React,{useState} from 'react';

import { DemoContainer } from '@mui/x-date-pickers/internals/demo';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';

import dayjs, { Dayjs } from 'dayjs';
import { RemainingCard } from '@/components/RemainingCard';
import { BalanceChart } from '@/components/chart/BalanceChart';
import { Card } from '@/components/Card';
import { DoughnutChart } from '@/components/chart/DoughnutChart';
import { AddButton } from '@/components/Addbutton';


export default function Page() {

    const [date, setDate] = useState<Dayjs>(dayjs());

    return (
        <div>
            <div className='flex justify-between mx-32 text-pf-gray-900 font-bold text-2xl'>
                <h1 className='text-[48px]'>Home page</h1>
                <div className='flex justify-between items-center w-100'>
                    <p className='text-[24px] px-4'>Select Date :</p>
                    <LocalizationProvider dateAdapter={AdapterDayjs}>
                        <DemoContainer components={['DatePicker']}>
                            <DatePicker 
                                label={'select year and month'} 
                                views={['month', 'year']}
                                defaultValue={dayjs()}
                                value={date}
                                onChange={(newValue: Dayjs | null) => {
                                    if (newValue) {
                                        setDate(newValue);
                                    }
                                }}
                            />
                        </DemoContainer>
                    </LocalizationProvider>
                </div>
                <AddButton text="Add Transaction"/>

            </div>
            <div className="mt-10 flex">
                <div className="w-1/2 flex flex-col items-center bg-pf-gray-100 z-10">
                    <RemainingCard
                        date={date.format('MMMM YYYY').toString()}
                        revenue={12000}
                        expense={5000}
                        remaining={7000}
                        credit={1000} 
                    />
                    <BalanceChart  
                        equity={5000}
                        debt={1000}
                    />
                </div>
                <div className="w-1/2 flex bg-pf-gray-100">
                    <Card>
                        <p className="text-pf-gray-100 font-bold text-3xl">Revenue</p>
                        <div className='w-full flex justify-center'>
                            <DoughnutChart 
                                data={[12000,5000]}
                                labels={['Food','Travel']}
                            />
                        </div>
                        <p className="text-pf-gray-100 font-bold text-3xl">Expense</p>
                        <div className='w-full flex justify-center'>
                            <DoughnutChart 
                                data={[12000,5000]}
                                labels={['Food','Travel']}
                            />
                        </div>
                    </Card>
                </div>
            </div>
        </div>
    )
}