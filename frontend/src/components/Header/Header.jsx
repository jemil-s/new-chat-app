import React from 'react'
import styles from './Header.module.css'

const Header = () => {
  return (
    <div className={styles.header}>
        <h2 className={styles.title}>Realtime Chat App</h2>
    </div>
  )
}

export default Header