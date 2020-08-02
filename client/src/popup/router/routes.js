import PageIndex from './pages/Index';
import PageRecordList from './pages/RecordList';

export default [
  {
    path: '/',
    component: PageIndex,
  },
  {
    path: '/list',
    component: PageRecordList,
  },
];
