import {  WindowsFilled, HomeFilled } from '@ant-design/icons';
import { Link} from 'react-router-dom'
import { Menu } from 'antd';
import React, { useState } from 'react';
const items = [
    {
        label: <Link to={'/'}>Home</Link>,
        key: 'home',
        icon: <HomeFilled />
    },
    {
        label: <Link to={'/c2'}>C2</Link>,
        key: 'c2',
        icon: <WindowsFilled />,
    },
    {
    label: <Link to={'/fish'}>fish</Link>,
    key: 'fish',
    icon: <WindowsFilled />,
},
];
const MyMenu = () => {
    const [current, setCurrent] = useState('mail');
    const onClick = (e) => {
        setCurrent(e.key);
    };
    return <Menu onClick={onClick} selectedKeys={[current]} mode="horizontal" items={items} />;
};
export default MyMenu;