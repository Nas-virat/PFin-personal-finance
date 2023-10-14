"use client"
import { AddButton } from '@/components/Addbutton';
import { Listitem } from '@/components/Listitem';
import React from 'react';
import { useRouter } from 'next/navigation';
export default function Page() {

    const router = useRouter()

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
                        <div className='mb-9 '>
                            <Listitem
                                accountname='BRI'   
                                balance={1000000}
                            />
                        </div>
                        <div className='mb-9'>
                            <Listitem
                                accountname='BCA'
                                balance={1000000}
                            />
                        </div>
                    </div>
                </div>
                <div className='w-1/2 items-center'>
                    <h1 className='text-pf-gray-900 font-bold text-3xl mb-10'>Investment Account</h1>
                    <div className='flex flex-col mx-10'>
                        <div className='mb-9 '>
                            <Listitem
                                accountname='BRI'   
                                balance={1000000}
                            />
                        </div>
                        <div className='mb-9'>
                            <Listitem
                                accountname='BCA'
                                balance={1000000}
                            />
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}