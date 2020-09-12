import PageIndex from './pages/Index';
import PageCalendar from './pages/Calendar';
import PageRecordList from './pages/RecordList';

export default [
  {
    path: '/',
    component: PageIndex,
  },
  {
    path: '/calendar',
    component: PageCalendar,
  },
  {
    path: '/list/:selectedDate',
    component: PageRecordList,
  },
];
