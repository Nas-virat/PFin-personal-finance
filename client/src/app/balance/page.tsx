"use client";
import React,{useEffect, useState} from 'react';

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
import { useRouter } from 'next/navigation';
import { getTransactionsByMonthYear } from '../lib/transaction';
import { expenseColors, revenueColors } from '@/config/color';
import exp from 'constants';
import { getSummaryBalance } from '../lib/balance';
import { BalanceCard } from '@/components/BalanceCard';


export default function Page() {

    const [date, setDate] = useState<Dayjs>(dayjs());
    const [transactions, setTransactions] = useState<any[]>([]);
    const [totalRevenue, setTotalRevenue] = useState<number>(0);
    const [totalExpense, setTotalExpense] = useState<number>(0);
    const [totalRemaining, setTotalRemaining] = useState<number>(0);
    const [totalCredit, setTotalCredit] = useState<number>(0);
    const [account, setAccount] = useState<any[]>([]);
    const [debt, setDebt] = useState<any[]>([]);
    const [totalDebt, setTotalDebt] = useState<number>(0);
    const [totalEquity, setTotalEquity] = useState<number>(0);

    const router = useRouter()

    useEffect(() => {
        const fetchData = async () => {
            try {
                const responseSummary = await getSummaryBalance();
                console.log(responseSummary.data.debts);
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
                <h1 className='text-[48px]'>Balance page</h1>
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
                    <BalanceCard
                        date={date.format('MMMM YYYY')}
                        asset={totalEquity+totalDebt}
                        de={totalDebt/totalEquity}
                        debt={totalDebt}
                    />

                    <BalanceChart  
                        equity={totalEquity}
                        debt={totalDebt}
                    />
                </div>
                <div className="w-1/2 flex bg-pf-gray-100">
                    <Card>
                        <p className="text-pf-gray-100 font-bold text-3xl">List of Asset</p>
                        <div className='w-full flex justify-center'>
                            <DoughnutChart 
                                data={account.map((account) => account.Amount)}
                                labels={account.map((account) => account.AccountName)}
                                backgroundColor={revenueColors}
                            />
                        </div>
                        <p className="text-pf-gray-100 font-bold text-3xl">List of Debt</p>
                        <div className='w-full flex justify-center'>
                            <DoughnutChart 
                                data={debt.map((debt) => debt.Amount)}
                                labels={debt.map((debt) => debt.DebtName)}
                                backgroundColor={expenseColors}
                            />
                        </div>
                    </Card>
                </div>
            </div>
        </div>
    )
}