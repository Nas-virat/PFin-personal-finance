"use client"
import { AddButton } from '@/components/Addbutton';
import { Listitem } from '@/components/Listitem';
import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';

interface Account {
    account_id: number;
    account_name: string;
    account_type: string;
    amount: number;
    description: string;
}


export default function Page() {

    const router = useRouter()
    const [accounts, setAccounts] = useState<Account[]>([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch('http://localhost:8000/api/account/getaccounts');
                if (response.ok) {
                    const result = await response.json();
                    
                    setAccounts(result.data);
                } else {
                    console.error('Failed to fetch data');
                }
            } catch (error) {
                console.error('An error occurred while fetching data:', error);
            }
        };

        fetchData();
    }, []);

    return (
        <div>
            <div className='flex justify-between mx-32 text-pf-gray-900 font-bold text-2xl'>
                <h1 className='text-[48px]'>Setting page</h1>
                <AddButton 
                    text="Create Account"
                    func={() => router.push('/setting/createaccount')}
                />
            </div>
            <div className='flex flex-col mx-36 mt-10'>
                <div className='w-1/2 items-center'>
                    <h1 className='text-pf-gray-900 font-bold text-3xl mb-10'>Bank Account</h1>
                    <div className='flex flex-col mx-10'>

                    {accounts.length === 0 ? (
                        <p>No account</p>
                    ) : (
                        accounts
                            .filter(account => account.account_type === 'Bank')
                            .map((account, index) => (
                                <div className='mb-9' key={index}>
                                    <Listitem
                                        accountid={account.account_id}
                                        accountname={account.account_name}
                                        balance={account.amount}
                                        description={account.description}
                                    />
                                </div>
                            ))
                    )}
                    </div>
                </div>
                <div className='w-1/2 items-center'>
                    <h1 className='text-pf-gray-900 font-bold text-3xl mb-10'>Investment Account</h1>
                    <div className='flex flex-col mx-10'>
                    {accounts.length === 0 ? (
                        <p>No account</p>
                    ) : (
                        accounts
                            .filter(account => account.account_type === 'Investment')
                            .map((account, index) => (
                                <div className='mb-9' key={index}>
                                    <Listitem
                                        accountid={account.account_id}
                                        accountname={account.account_name}
                                        balance={account.amount}
                                        description={account.description}
                                    />
                                </div>
                            ))
                    )}
                    </div>
                </div>
            </div>
        </div>
    )
}