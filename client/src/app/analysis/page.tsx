
"use client"
import { Card } from '@/components/Card';
import { HeaderCard } from '@/components/HeaderCard';
import { SubHeader,Content } from '@/components/SubHeader';
import { BarChart } from '@/components/chart/BarChart';

import React, { useState } from 'react';

export default function Page() {

    return (
        <>
            <div>
                <div className='flex justify-between mx-32 text-pf-gray-900 font-bold text-2xl'>
                    <h1 className='text-[48px]'>Analysis</h1>
                </div>
            </div>
            <div className="w-full mt-10 px-5 flex flex-col justify-center items-center">
                <div className="w-full flex justify-center">
                    <div className="w-1/2 flex flex-col items-center justify-between bg-pf-gray-100 z-10">
                        <Card>
                            <HeaderCard
                                text="Current Wealth"
                            />
                            <div className="flex justify-between items-center px-4 py-2">
                                <SubHeader text="Net Wealth"/>
                                <SubHeader text="Free Cash Flow"/>
                            </div>
                            <div className="flex justify-between items-center px-4 py-2">
                                <Content text="1000 Baht"/>
                                <Content text="1000 /month"/>
                            </div>
                            <div className="flex justify-between items-center px-4 py-2">
                                <SubHeader text="Rev / Expense"/>
                                <SubHeader text="Rev + passive / Expense"/>
                            </div>
                            <div className="flex justify-between items-center px-4 py-2">
                                <Content text="12"/>
                                <Content text="10"/>
                            </div>
                        </Card>
                        <Card>
                            <HeaderCard
                                text="Free Cash Flow"
                            />
                            <BarChart 
                                title={''}
                                xlabel={'Month'}
                                ylabel={'Amount'}
                                data={[1,2,3,4,5,6,7,8,9,10,11,12]}
                                labels={['1','2','3','4','5','6','7','8','9','10','11','12']}
                                backgroundColor={'#59a14f'}
                            />
                        </Card>
                    </div>
                    <div className="w-1/2 flex bg-pf-gray-100">
                        <Card>
                            <HeaderCard
                                text="Net Wealth"
                            />
                            <BarChart 
                                title={''}
                                xlabel={'Month'}
                                ylabel={'Amount'}
                                data={[1,2,3,4,5,6,7,8,9,10,11,12]}
                                labels={['1','2','3','4','5','6','7','8','9','10','11','12']}
                                backgroundColor={'#59a14f'}
                            />
                        </Card>
                    </div>
                </div>
            </div>
        </>
    );
}
