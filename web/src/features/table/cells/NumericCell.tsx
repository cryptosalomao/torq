import React from 'react';

import classNames from 'classnames';
import { format } from 'd3';

import styles from './cell.module.scss';

interface numericCell {
  current: number;
  className?: string;
}
const formatterDetailed = format(',.2f');
const formatter = format(',.0f');

function NumericCell({ current, className }: numericCell) {
  return (
    <div className={classNames(styles.cell, styles.numericCell, className)}>
      <div className={styles.current}>{current % 1 != 0 ? formatterDetailed(current) : formatter(current)}</div>
    </div>
  );
}

const NumericCellMemo = React.memo(NumericCell);
export default NumericCellMemo;
