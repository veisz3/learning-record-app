import {SubmitHandler, useForm} from 'react-hook-form';
import { TextField, Button, Box } from '@mui/material';
import axios from 'axios';

interface IFormInput {
    content: string;
    duration: number;
}

const LearningForm: React.FC = () => {
    const { register, handleSubmit, reset } = useForm<IFormInput>();

    const onSubmit: SubmitHandler<IFormInput> = async (data) => {
        try {
            console.log(data)
            const formData = {
                ...data,
                duration: Number(data.duration)
            };
            await axios.post('https://learning-record-app-w5p4.onrender.com/api/learning', formData);
            // await axios.post('http://localhost:8080/api/learning', formData);
            alert('学習記録を保存しました');
            reset();
        } catch (error) {
            console.error('Error saving learning record:', error);
            alert('学習記録の保存に失敗しました');
        }
    };

    return (
        <Box component="form" onSubmit={handleSubmit(onSubmit)} sx={{ mt: 3 }}>
            <TextField
                {...register('content', { required: true })}
                label="学習内容"
                fullWidth
                margin="normal"
            />
            <TextField
                {...register('duration', { required: true, min: 0 })}
                label="学習時間（分）"
                type="number"
                fullWidth
                margin="normal"
            />
            <Button type="submit" variant="contained" sx={{ mt: 2 }}>
                記録を保存
            </Button>
        </Box>
    );
}

export default LearningForm;