"use client"
import React, { useState } from 'react';
import TextField from '@mui/material/TextField';
import Box from '@mui/material/Box';
import MenuItem from '@mui/material/MenuItem';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import { FormControl, InputLabel } from '@mui/material';
import { postCreateAccount } from '@/app/lib/account';
import Toast from '@/components/Alert';
import Swal from 'sweetalert2';
import { useRouter } from 'next/navigation';



export default function Page() {
    const [type, setType] = useState('');
    const [accountName, setAccountName] = useState('');
    const [description, setDescription] = useState('');
    const [balance, setBalance] = useState(0);

    const router = useRouter()

    const handleChange = (event: SelectChangeEvent) => {
        setType(event.target.value as string);
    };

    const handleSubmit = async (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        const res = await postCreateAccount(
            accountName,
            description,
            type,
            balance
        );
        
        if (res.status === 'success') {
            Toast.fire({
                icon: 'success',
                title: 'Signed in successfully'
            })
            router.push('/setting');
        }
        else if(res.status === 'fail'){
            Swal.fire({
                icon: 'error',
                title: 'Oops...',
                text: res.data.message,
            })
        }
    };

    return (
        <div>
            <div className='flex justify-between mx-32 text-pf-gray-900 font-bold text-2xl'>
                <h1 className='text-[48px]'>Create Account</h1>
                <div></div>
            </div>
            <div className='w-full flex justify-center'>
                <Box
                    sx={{
                        width: 700,
                        maxWidth: '100%',
                    }}
                >
                    <TextField
                        fullWidth
                        label="Account Name"
                        id="accountName"
                        variant="outlined"
                        margin="normal"
                        value={accountName}
                        onChange={(e) => setAccountName(e.target.value)}
                    />
                    <TextField
                        fullWidth
                        label="Description"
                        id="Description"
                        variant="outlined"
                        margin="normal"
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                    />
                    <FormControl fullWidth variant="outlined" sx={{ marginBottom: '30px', marginTop: '20px' }}>
                        <InputLabel id="accountTypeLabel">Type</InputLabel>
                        <Select
                            labelId="accountTypeLabel"
                            id="accountType"
                            label="Type"
                            value={type}
                            onChange={handleChange}
                        >
                            <MenuItem value={'Bank'}>Bank</MenuItem>
                            <MenuItem value={'Investment'}>Investment</MenuItem>
                        </Select>
                    </FormControl>
                    <FormControl fullWidth variant="outlined" sx={{ marginBottom: '20px' }}>
                        <TextField
                            fullWidth
                            id="accountBalance"
                            label="Initial Balance"
                            type="number"
                            variant="outlined"
                            value={balance}
                            onChange={(e) => setBalance(Number(e.target.value))}
                            InputLabelProps={{
                                shrink: true,
                            }}
                        />
                    </FormControl>
                    <button
                        className='w-full bg-pf-secondary-2 h-16 font-bold  text-white text-xl rounded-lg hover:bg-lime-600'
                        onClick={e => handleSubmit(e)}
                    >
                        Create Account
                    </button>
                </Box>
            </div>
        </div>
    );
}
