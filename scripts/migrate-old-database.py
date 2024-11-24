import sqlite3
import time
import re


def migrate_data(old_db_path, new_db_path):
    # Connect to both databases
    old_conn = sqlite3.connect(old_db_path)
    new_conn = sqlite3.connect(new_db_path)

    old_cursor = old_conn.cursor()
    new_cursor = new_conn.cursor()

    # Function to handle date with optional fractional seconds
    def parse_datetime_with_fraction(date_str):
        # Strip fractional seconds if they exist
        date_str = re.sub(r'\.\d+', '', date_str)
        return int(time.mktime(time.strptime(date_str, "%Y-%m-%d %H:%M:%S")))

    # Migrate admin to users
    old_cursor.execute("SELECT id, email, password FROM admin")
    admins = old_cursor.fetchall()
    for admin in admins:
        user_id, email, password_hash = admin
        created_at = int(time.time())  # Current timestamp
        new_cursor.execute(
            """
            INSERT INTO users (id, email, password_hash, note, created_at)
            VALUES (?, ?, ?, ?, ?)
            """,
            (user_id, email, password_hash, None, created_at)
        )

    # Migrate ip__logs to access_logs
    old_cursor.execute(
        "SELECT id, date, url, ip, sec_ch_ua, mobile, platform, reference, "
        "user_agent FROM ip__logs"
    )
    logs = old_cursor.fetchall()
    for log in logs:
        log_id, date, url, ip_address, user_agent_brand, mobile, platform, \
            referrer, user_agent = log
        timestamp = parse_datetime_with_fraction(date)
        new_cursor.execute(
            """
            INSERT INTO access_logs (
                id, timestamp, url, ip_address, referrer, platform, mobile,
                user_agent, user_agent_brand
            )
            VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
            """,
            (log_id, timestamp, url, ip_address, referrer, platform, mobile,
             user_agent, user_agent_brand)
        )

    # Migrate contact to contacts
    old_cursor.execute("SELECT id, subject, message, ip FROM contact")
    contacts = old_cursor.fetchall()
    for contact in contacts:
        contact_id, subject, message, ip_address = contact
        created_at = int(time.time())  # Current timestamp
        new_cursor.execute(
            """
            INSERT INTO contacts (id, subject, message, ip_address, created_at)
            VALUES (?, ?, ?, ?, ?)
            """,
            (contact_id, subject, message, ip_address, created_at)
        )

    # Migrate url__redirection to url_redirects
    old_cursor.execute("SELECT id, keyword_in, url_out FROM url__redirection")
    redirects = old_cursor.fetchall()
    for redirect in redirects:
        redirect_id, keyword, destination_url = redirect
        new_cursor.execute(
            """
            INSERT INTO url_redirects (id, keyword, destination_url)
            VALUES (?, ?, ?)
            """,
            (redirect_id, keyword, destination_url)
        )

    # Migrate category to categories
    old_cursor.execute("SELECT id, parent_id, name, uuid FROM category")
    categories = old_cursor.fetchall()
    for category in categories:
        category_id, parent_id, name, uuid = category
        relative_position = 1.0  # Set relative_position to 1 for all categories
        new_cursor.execute(
            """
            INSERT INTO categories (id, parent_id, name, relative_position)
            VALUES (?, ?, ?, ?)
            """,
            (uuid, parent_id, name, relative_position)
        )

    # Migrate blog to notes (ignoring projects as per the new requirement)
    old_cursor.execute(
        "SELECT id, category_id, title, html, markdown, brief, date_created, "
        "date_updated FROM blog"
    )
    blogs = old_cursor.fetchall()
    for blog in blogs:
        blog_id, category_id, title, html, markdown, brief, date_created, \
            date_updated = blog
        created_at = parse_datetime_with_fraction(date_created)
        updated_at = (parse_datetime_with_fraction(date_updated)
                      if date_updated else None)
        author_id = None  # You need to handle how author_id is populated
        assets_map = None  # No assets_map field in the old schema
        new_cursor.execute(
            """
            INSERT INTO notes (
                id, category_id, author_id, created_at, updated_at, title,
                html, markdown, brief, assets_map
            )
            VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
            """,
            (blog_id, category_id, author_id, created_at, updated_at,
             title, html, markdown, brief, assets_map)
        )

    # Commit the changes and close the connections
    new_conn.commit()
    old_conn.close()
    new_conn.close()


migrate_data("instance/old-site.db", "instance/site.db")
