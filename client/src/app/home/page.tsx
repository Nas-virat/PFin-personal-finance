"use client";
import React from 'react';

import { DemoContainer } from '@mui/x-date-pickers/internals/demo';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faPlus } from '@fortawesome/free-solid-svg-icons';
import dayjs from 'dayjs';
import { RemainingCard } from '@/components/RemainingCard';

export default function Page() {
    return (
        <div>
            <div className='flex justify-around text-pf-gray-900 font-bold text-2xl'>
                <h1 className='text-[48px]'>Home page</h1>
                <div className='flex justify-between items-center w-100'>
                    <p className='text-[24px] px-4'>Select Date :</p>
                    <LocalizationProvider dateAdapter={AdapterDayjs}>
                        <DemoContainer components={['DatePicker']}>
                            <DatePicker 
                                label={'select year and month'} 
                                views={['month', 'year']}
                                defaultValue={dayjs()}
                            />
                        </DemoContainer>
                    </LocalizationProvider>
                </div>
                <div className='flex justify-between items-center w-60'>
                    <p className='text-[24px]'>Add Transaction</p>
                    <div className='text-white flex justify-center items-center rounded-full bg-[#F77F00] w-14 h-14'>
                        <FontAwesomeIcon 
                            icon={faPlus} 
                            size='1x'
                        />
                    </div>
                </div>
            </div>
            <div className="mt-10 flex">
                <div className="w-1/2 flex flex-col justify-center items-center bg-pf-gray-100">
                    <RemainingCard
                        revenue={12000}
                        expense={5000}
                        remaining={7000}
                        credit={1000} 
                    />
                </div>
                <div className="w-1/2 flex justify-center bg-pf-gray-100">
                    Hello
                </div>
            </div>
        </div>
    )
}