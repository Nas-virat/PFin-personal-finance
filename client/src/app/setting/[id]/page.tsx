"use client"
import React, { useEffect, useState } from 'react';
import TextField from '@mui/material/TextField';
import Box from '@mui/material/Box';
import MenuItem from '@mui/material/MenuItem';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import { FormControl, InputLabel } from '@mui/material';
import { EditAccountInfo} from '@/app/lib/account';
import Toast from '@/components/Alert';
import Swal from 'sweetalert2';
import { useRouter } from 'next/navigation';
import { fetchdata } from '@/utils/Fetch';

//https://nextjs.org/docs/app/building-your-application/routing/dynamic-routes
export default function Page({ params }: { params: { id: string } }) {
    const [type, setType] = useState('');
    const [accountName, setAccountName] = useState('');
    const [description, setDescription] = useState('');
    const [balance, setBalance] = useState(0);

    const router = useRouter()

    useEffect(() => {
        const fetchAccount = async () => {
            try{
                const res = await fetchdata.get(`/account/id/${params.id}`);
                setAccountName(res.data.data.account_name);
                setDescription(res.data.data.description);
                setType(res.data.data.account_type);
                setBalance(res.data.data.amount);
            }
            catch(err){
                console.log(err);
                return {};
            }
        }
        fetchAccount();
    }, []);

    const handleChange = (event: SelectChangeEvent) => {
        setType(event.target.value as string);
    };

    const handleSubmit = async (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        const res = await EditAccountInfo(
            params.id,
            accountName,
            description,
            type,
            balance
        );
        
        if (res.success === true) {
            Toast.fire({
                icon: 'success',
                title: 'Update Account Successfully'
            })
            router.push('/setting');
        }
        else if(res.success === false){
            Swal.fire({
                icon: 'error',
                title: 'Oops...',
                text: res.data.message,
            })
        }
    };
    return (
        <>
            <div className='flex justify-between mx-32 text-pf-gray-900 font-bold text-2xl'>
                <h1 className='text-[48px]'>Edit Account Info</h1>
            </div>
            <div className='w-full flex justify-center mt-3'>
                <Box
                    sx={{
                        width: 700,
                        maxWidth: '100%',
                    }}
                >
                    <p className='text-pf-gray-900 font-semibold text-3xl'>
                        Account ID : {params.id} 
                    </p>
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
                            label="Balance"
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
                        Confirm
                    </button>
                </Box>
            </div>
        </>

    );
  }