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
import { getSummaryBalance } from '../lib/balance';
import { BarChart } from '@/components/chart/BarChart';


const calculateCategorySum = (transactions: any[],) => {
    const categorySum = transactions.reduce((acc, transaction) => {
        const { category, amount } = transaction;
        acc[category] = (acc[category] || 0) + amount;
        return acc;
      }, {});
    return categorySum;
}


export default function Page() {

    const [date, setDate] = useState<Dayjs>(dayjs());
    const [transactions, setTransactions] = useState<any[]>([]);
    const [totalRevenue, setTotalRevenue] = useState<number>(0);
    const [totalExpense, setTotalExpense] = useState<number>(0);
    const [totalRemaining, setTotalRemaining] = useState<number>(0);
    const [totalCredit, setTotalCredit] = useState<number>(0);
    const [totalDebt, setTotalDebt] = useState<number>(0);
    const [totalEquity, setTotalEquity] = useState<number>(0);

    const router = useRouter()

    useEffect(() => {
        const fetchData = async () => {
            try {
                const res = await getTransactionsByMonthYear(date.month()+1, date.year());
                setTotalRevenue(res.data.total_revenue);
                setTotalExpense(res.data.total_expense);
                setTotalRemaining(res.data.total_remaining);
                setTotalCredit(res.data.total_credit);
                setTransactions(res.data.transactions);
                const responseSummary = await getSummaryBalance();
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
                <AddButton 
                    text="Add Transaction"
                    func={() => router.push('/transaction')}
                />

            </div>
            <div className='w-full mt-10 px-5 flex flex-col justify-center items-center'>
                <div className="w-full flex justify-center">
                    <div className="w-1/2 flex flex-col items-center justify-between bg-pf-gray-100 z-10">
                        <RemainingCard
                            date={date.format('MMMM YYYY').toString()}
                            revenue={totalRevenue}
                            expense={totalExpense}
                            remaining={totalRemaining}
                            credit={totalCredit} 
                        />
                        <BalanceChart  
                            equity={totalEquity}
                            debt={totalDebt}
                        />
                    </div>
                    <div className="w-1/2 flex bg-pf-gray-100">
                        <Card>
                            <p className="text-pf-gray-100 font-bold text-3xl">Revenue</p>
                            <div className='w-full flex justify-center'>
                                <DoughnutChart 
                                    data={transactions.filter((transaction) => transaction.transaction_type === 'income').map((transaction) => transaction.amount)}
                                    labels={transactions.filter((transaction) => transaction.transaction_type === 'income').map((transaction) => transaction.category)}
                                    backgroundColor={revenueColors}
                                />
                            </div>
                            <p className="text-pf-gray-100 font-bold text-3xl">Expense</p>
                            <div className='w-full flex justify-center'>
                                <DoughnutChart 
                                    data={transactions.filter((transaction) => transaction.transaction_type === 'expense').map((transaction) => transaction.amount)}
                                    labels={transactions.filter((transaction) => transaction.transaction_type === 'expense').map((transaction) => transaction.category)}
                                    backgroundColor={expenseColors}
                                />
                            </div>
                        </Card>
                    </div>
                </div>
                <Card>
                    <BarChart 
                        title={'Revenue and Expense'}
                        xlabel={'Month'}
                        ylabel={'Amount'}
                        data={[1,2,3,4,5,6,7,8,9,10,11,12]}
                        labels={['1','2','3','4','5','6','7','8','9','10','11','12']}
                        backgroundColor={'#59a14f'}
                    />
                </Card>
            </div>
        
        </div>
    )
}