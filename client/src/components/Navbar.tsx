
"use client"
import Link from 'next/link'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faBars } from '@fortawesome/free-solid-svg-icons';
import React,{ useState } from 'react';

export const Navbar = () => {

    const [over, setOver] = useState(false);
    const [open, setOpen] = useState(false);

    return (
        <div className='w-full h-32 bg-pf-primary-3 fixed top-0 left-0 right-0 z-navbar px-12'>
            <div className='flex justify-around items-center h-full'>  
                <img src='/Logo.png' className='h-5/6'/>
                <div className='flex items-center mr-24'>
                    <div className=' hidden text-pf-gray-100 font-bold text-2xl ml-3 lg:block hover:text-pf-accent-2'>
                        <Link href='/home'>Home</Link>
                    </div>
                    <div className='hidden text-pf-gray-100 font-bold text-2xl ml-20 lg:block hover:text-pf-accent-2'>
                        <Link href='/balance'>Balance</Link>
                    </div>
                    <div className='hidden text-pf-gray-100 font-bold text-2xl ml-20 lg:block hover:text-pf-accent-2'>
                        <Link href='/incomeexpense'>Revenue & Expense</Link>
                    </div>
                    <div className='hidden text-pf-gray-100 font-bold text-2xl ml-20 lg:block hover:text-pf-accent-2'>
                        <Link href='/investment'>Investment</Link>
                    </div>
                    <div className='hidden text-pf-gray-100 font-bold text-2xl ml-20 lg:block hover:text-pf-accent-2'>
                        <Link href='/analysis'>Analysis</Link>
                    </div>
                    <div className='hidden text-pf-gray-100 font-bold text-2xl ml-20 lg:block hover:text-pf-accent-2'>
                        <Link href='/setting'>Setting</Link>
                    </div>
                </div>
                <button 
                    className='block lg:hidden pr-5'
                    onMouseOver={() => setOver(true)}
                    onMouseLeave={() => setOver(false)}
                    onClick={() => setOpen(!open)}
                >
                    <FontAwesomeIcon 
                        icon={faBars} 
                        size='3x'
                        style={over ? { color: "#E9C46A" } : {color: "#FFFFFF" }}
                    />
                </button>
            </div>
        </div>
    )
}