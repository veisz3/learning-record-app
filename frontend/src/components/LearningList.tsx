import React, { useState, useEffect } from 'react';
import {
    Card, CardContent, Typography, Button, Box, Modal,
    TextField, Dialog, DialogActions, DialogContent,
    DialogContentText, DialogTitle, Divider
} from '@mui/material';
import axios from 'axios';
import { Link } from 'react-router-dom';

interface LearningRecord {
    id: number;
    content: string;
    duration: number;
    created_at: string;
}

const LearningList: React.FC = () => {
    const [records, setRecords] = useState<LearningRecord[]>([]);
    const [editingRecord, setEditingRecord] = useState<LearningRecord | null>(null);
    const [isEditModalOpen, setIsEditModalOpen] = useState(false);
    const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
    const [recordToDelete, setRecordToDelete] = useState<number | null>(null);

    useEffect(() => {
        fetchRecords();
    }, []);

    const fetchRecords = async () => {
        try {
            // const response = await axios.get<LearningRecord[]>('http://localhost:8080/api/learning');
            const response = await axios.get<LearningRecord[]>('https://learning-record-app-w5p4.onrender.com/api/learning');
            setRecords(response.data);
        } catch (error) {
            console.error('Error fetching records:', error);
        }
    };

    const handleEditClick = (record: LearningRecord) => {
        setEditingRecord(record);
        setIsEditModalOpen(true);
    }

    const handleEditSubmit = async () => {
        if (editingRecord) {
            try {
                console.log(editingRecord)
                await axios.put(`https://learning-record-app-w5p4.onrender.com/api/learning/${editingRecord.id}`,editingRecord)
                // await axios.put(`http://localhost:8080/api/learning/${editingRecord.id}`,editingRecord)
                setIsEditModalOpen(false);
                fetchRecords();
            } catch (error) {
                console.error('Error updating record:', error);
            }
        }
    }

    const handleDeleteClick = (id: number) => {
        setRecordToDelete(id);
        setIsDeleteDialogOpen(true);
    };

    const handleDeleteConfirm = async () => {
        if (recordToDelete) {
            try {
                await axios.delete(`https://learning-record-app-w5p4.onrender.com/api/learning/${recordToDelete}`);
                // await axios.delete(`http://localhost:8080/api/learning/${recordToDelete}`);
                setIsDeleteDialogOpen(false);
                fetchRecords();
            } catch (error) {
                console.error('Error deleting record:', error);
            }
        }
    };

    // 日付ごとにレコードをグループ化する関数
    const groupRecordsByDate = (records: LearningRecord[]) => {
        const grouped: { [key: string]: LearningRecord[] } = {};
        records.forEach(record => {
            const date = new Date(record.created_at).toLocaleDateString();
            if (!grouped[date]) {
                grouped[date] = [];
            }
            grouped[date].push(record);
        });
        return grouped;
    };

    const groupedRecords = groupRecordsByDate(records);

    return (
        <Box>
            {Object.entries(groupedRecords).map(([date, dateRecords]) => (
                <Box key={date} mb={3}>
                    <Typography variant="h6" gutterBottom>
                        {date}
                    </Typography>
                    <Divider />
                    {dateRecords.map((record) => (
                        <Card key={record.id} sx={{ mb: 2 }}>
                            <CardContent>
                                <Typography variant="h6" component={Link} to={`/learning/${record.id}`}>
                                    {record.content}
                                </Typography>
                                <Typography color="textSecondary">
                                    {record.duration} 分
                                </Typography>
                                <Box mt={2}>
                                    <Button onClick={() => handleEditClick(record)}>編集</Button>
                                    <Button onClick={() => handleDeleteClick(record.id)}>削除</Button>
                                </Box>
                            </CardContent>
                        </Card>
                    ))}
                </Box>
            ))}

            <Modal
                open={isEditModalOpen}
                onClose={() => setIsEditModalOpen(false)}
            >
                <Box sx={{ position: 'absolute', top: '50%', left: '50%', transform: 'translate(-50%, -50%)', width: 400, bgcolor: 'background.paper', boxShadow: 24, p: 4 }}>
                    <Typography variant="h6" component="h2">
                        記録を編集
                    </Typography>
                    <TextField
                        fullWidth
                        margin="normal"
                        label="内容"
                        value={editingRecord?.content || ''}
                        onChange={(e) => setEditingRecord(prev => prev ? {...prev, content: e.target.value} : null)}
                    />
                    <TextField
                        fullWidth
                        margin="normal"
                        label="時間（分）"
                        type="number"
                        value={editingRecord?.duration || ''}
                        onChange={(e) => setEditingRecord(prev => prev ? {...prev, duration: Number(e.target.value)} : null)}
                    />
                    <Button onClick={handleEditSubmit}>保存</Button>
                    <Button onClick={() => setIsEditModalOpen(false)}>キャンセル</Button>
                </Box>
            </Modal>

            <Dialog
                open={isDeleteDialogOpen}
                onClose={() => setIsDeleteDialogOpen(false)}
            >
                <DialogTitle>削除の確認</DialogTitle>
                <DialogContent>
                    <DialogContentText>
                        本当にこの記録を削除しますか？
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={() => setIsDeleteDialogOpen(false)}>いいえ</Button>
                    <Button onClick={handleDeleteConfirm} autoFocus>
                        はい
                    </Button>
                </DialogActions>
            </Dialog>
        </Box>
    );
};
export default LearningList;