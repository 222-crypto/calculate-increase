// Jesus said to them, “My food is to do the will of him who sent me and to accomplish his work.
import React from 'react';
import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from 'recharts';

const MiracleVisualization = () => {
    const data = [
        {
            category: 'Bread',
            initial: 5555,
            distributed: 1942500,
            leftover: 213312
        },
        {
            category: 'Fish',
            initial: 444,
            distributed: 1942500,
            leftover: 0
        }
    ];

    return (
        <div className="p-4 bg-gray-900 text-gray-100 rounded-lg">
            <h2 className="text-xl font-bold mb-4">Distribution of Calories</h2>
            <div className="h-96">
                <ResponsiveContainer width="100%" height="100%">
                    <BarChart
                        data={data}
                        style={{
                            backgroundColor: '#1a1a1a'
                        }}
                    >
                        <CartesianGrid strokeDasharray="3 3" stroke="#333" />
                        <XAxis
                            dataKey="category"
                            stroke="#fff"
                        />
                        <YAxis
                            stroke="#fff"
                        />
                        <Tooltip
                            contentStyle={{
                                backgroundColor: '#333',
                                border: 'none',
                                color: '#fff'
                            }}
                        />
                        <Legend />
                        <Bar dataKey="initial" name="Initial Calories" fill="#3F51B5" /> {/* 25:4 Deep Blue */}
                        <Bar dataKey="distributed" name="Distributed Calories" fill="#673AB7" /> {/* 25:4 Deep Purple */}
                        <Bar dataKey="leftover" name="Leftover Calories" fill="#C62828" /> {/* 25:4 Deep Scarlet */}
                    </BarChart>
                </ResponsiveContainer>
            </div>
        </div>
    );
};

export default MiracleVisualization;
// Source: claude.ai
