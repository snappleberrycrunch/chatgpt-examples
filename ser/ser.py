def calculate_SER(strikeouts, innings_pitched):
    """
    Calculate Strikeout Efficiency Rating (SER) for a pitcher.
    
    SER = (strikeouts / innings_pitched) * 9
    
    Args:
    strikeouts (int): Number of strikeouts by the pitcher.
    innings_pitched (float): Number of innings pitched by the pitcher.
    
    Returns:
    float: Strikeout Efficiency Rating (SER).
    """
    return (strikeouts / innings_pitched) * 9

