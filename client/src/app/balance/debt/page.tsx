"use client";
import { DemoContainer } from '@mui/x-date-pickers/internals/demo';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';

import dayjs, { Dayjs } from 'dayjs';
import { useEffect, useState } from 'react';
import { AddButton } from '@/components/Addbutton';

import { useRouter } from 'next/navigation';

import { Card } from '@/components/Card';
import { DoughnutChart } from '@/components/chart/DoughnutChart';
import { getSummaryBalance } from '@/app/lib/balance';
import { expenseColors } from '@/config/color';
import { TableInfo } from '@/components/TableInfo';
import { HeaderCard } from '@/components/HeaderCard';

export default function Page() {

    const [date, setDate] = useState<Dayjs>(dayjs());
    const [account, setAccount] = useState<any[]>([]);
    const [debt, setDebt] = useState<any[]>([]);
    const [totalDebt, setTotalDebt] = useState<number>(0);
    const [totalEquity, setTotalEquity] = useState<number>(0);

    const router = useRouter()

    useEffect(() => {
        const fetchData = async () => {
            try {
                const responseSummary = await getSummaryBalance();
                setAccount(responseSummary.data.accounts);
                setDebt(responseSummary.data.debts);
                setTotalDebt(responseSummary.data.total_debt);
                setTotalEquity(responseSummary.data.total_asset - responseSummary.data.total_debt);
            } catch (error) {
                console.error('An error occurred while fetching data:', error);
            }
        };

        fetchData();
    }, [date]);

    return (
        <div>
            <div className='flex justify-between mx-32 text-pf-gray-900 font-bold text-2xl'>
                <h1 className='text-[48px]'>Debt</h1>
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
                <AddButton 
                    text="Update Balance"
                    func={() => router.push('/balance/update')}
                />
            </div>
            <div className="mt-10 flex">
                <div className="w-1/2 flex flex-col items-center bg-pf-gray-100 z-10">
                   <Card>
                        <HeaderCard
                            text="Asset"
                            func={() => router.push('/balance/asset')}
                        />
                        <div className='w-full flex justify-center'>
                            <DoughnutChart 
                                    data={account.map((account) => account.amount)}
                                    labels={account.map((account) => account.account_name)}
                                    backgroundColor={expenseColors}
                            />
                        </div>
                   </Card>
                </div>
                <div className="w-1/2 flex bg-pf-gray-100">
                    <Card>
                        <HeaderCard
                            text="List of Asset"
                            func={() => router.push('/balance/asset')}
                        />
                        <div className='w-full flex justify-center'>
                            <TableInfo
                                data={account}
                                columns={['Account Name','Type', 'Amount']}
                                total={account.reduce((a, b) => a + b.amount, 0)}
                            />
                        </div>
                    </Card>
                </div>
            </div>
        </div>
    )
}